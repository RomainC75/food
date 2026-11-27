import { describe, expect, it, vi } from "vitest";
import {
  AxiosAdapter,
  AxiosError,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from "axios";
import { createHttpClient, REFRESH_URL } from "./http-client";
import { InMemTokenStorage } from "./in-mem-tokenStorage";
import { SessionExpiredError } from "../domain/errors";

type Route = (config: InternalAxiosRequestConfig) => { status: number; data?: unknown };

// Minimal fake backend: no network, no extra dependency
const fakeAdapter = (route: Route): AxiosAdapter => async (config) => {
  const { status, data } = route(config);
  const response: AxiosResponse = { status, statusText: String(status), data, headers: {}, config };
  if (status >= 400) {
    throw new AxiosError(`status ${status}`, "ERR_BAD_REQUEST", config, null, response);
  }
  return response;
};

const bearer = (config: InternalAxiosRequestConfig) =>
  String(config.headers.Authorization ?? "");

const setup = (route: Route) => {
  const tokenStorage = new InMemTokenStorage();
  tokenStorage.saveTokens("refresh-1", "access-old");
  const onSessionExpired = vi.fn();
  const adapter = vi.fn(fakeAdapter(route));
  const api = createHttpClient({ tokenStorage, onSessionExpired, axiosConfig: { adapter } });
  const calls = (url: string) => adapter.mock.calls.filter(([c]) => c.url === url);
  return { api, tokenStorage, onSessionExpired, calls };
};

describe("http client", () => {
  it("should attach the access token", async () => {
    const { api, calls } = setup(() => ({ status: 200 }));
    await api.get("/data");
    expect(bearer(calls("/data")[0][0])).toBe("Bearer access-old");
  });

  it("should not attach the token when skipAuth is set", async () => {
    const { api, calls } = setup(() => ({ status: 200 }));
    await api.post("/auth/login", {}, { skipAuth: true });
    expect(bearer(calls("/auth/login")[0][0])).toBe("");
  });

  it("should refresh on 401 and replay the request with the new token", async () => {
    const { api, tokenStorage, calls, onSessionExpired } = setup((config) => {
      if (config.url === REFRESH_URL) return { status: 200, data: { access_token: "access-new" } };
      return bearer(config) === "Bearer access-new"
        ? { status: 200, data: "ok" }
        : { status: 401 };
    });

    const res = await api.get("/data");

    expect(res.data).toBe("ok");
    expect(tokenStorage.getAccessToken()).toBe("access-new");
    expect(JSON.parse(calls(REFRESH_URL)[0][0].data)).toEqual({ refresh_token: "refresh-1" });
    expect(onSessionExpired).not.toHaveBeenCalled();
  });

  it("should share ONE refresh between parallel 401s", async () => {
    const { api, calls } = setup((config) => {
      if (config.url === REFRESH_URL) return { status: 200, data: { access_token: "access-new" } };
      return bearer(config) === "Bearer access-new" ? { status: 200 } : { status: 401 };
    });

    await Promise.all([api.get("/a"), api.get("/b"), api.get("/c")]);

    expect(calls(REFRESH_URL)).toHaveLength(1);
    expect(calls("/a")).toHaveLength(2); // first try + replay
  });

  it("should end the session when the refresh is rejected", async () => {
    const { api, tokenStorage, onSessionExpired } = setup((config) =>
      config.url === REFRESH_URL ? { status: 400 } : { status: 401 },
    );

    await expect(api.get("/data")).rejects.toBeInstanceOf(SessionExpiredError);
    expect(tokenStorage.hasSession()).toBe(false);
    expect(onSessionExpired).toHaveBeenCalledOnce();
  });

  it("should end the session when the replay is still 401", async () => {
    const { api, calls, onSessionExpired } = setup((config) =>
      config.url === REFRESH_URL
        ? { status: 200, data: { access_token: "access-new" } }
        : { status: 401 },
    );

    await expect(api.get("/data")).rejects.toBeInstanceOf(SessionExpiredError);
    expect(calls("/data")).toHaveLength(2); // no infinite loop
    expect(onSessionExpired).toHaveBeenCalledOnce();
  });

  it("should not refresh on a 401 from a skipAuth request (e.g. login)", async () => {
    const { api, calls, onSessionExpired } = setup(() => ({ status: 401 }));

    await expect(api.post("/auth/login", {}, { skipAuth: true })).rejects.toBeInstanceOf(AxiosError);
    expect(calls(REFRESH_URL)).toHaveLength(0);
    expect(onSessionExpired).not.toHaveBeenCalled();
  });

  it("should not touch the session on other errors", async () => {
    const { api, tokenStorage, onSessionExpired } = setup(() => ({ status: 500 }));

    await expect(api.get("/data")).rejects.toBeInstanceOf(AxiosError);
    expect(tokenStorage.hasSession()).toBe(true);
    expect(onSessionExpired).not.toHaveBeenCalled();
  });
});
