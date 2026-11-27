import { describe, expect, it, vi } from "vitest";
import { createTestDependencies, TestDependenciesOptions } from "@/dependencies";
import { setAuthStore } from "./auth.store";

const setup = (opts: TestDependenciesOptions = {}) => {
  const deps = createTestDependencies(opts);
  const store = setAuthStore(deps.authSrv, deps.workspaceSrv);
  return { deps, store };
};

describe("auth store", () => {
  describe("initialize", () => {
    it("should be unauthenticated without a session", async () => {
      const { store } = setup();
      await store.getState().initialize();
      expect(store.getState().status).toBe("unauthenticated");
    });

    it("should restore the user from a stored session", async () => {
      const { store, deps } = setup();
      deps.tokenStorage.saveTokens("refresh", "access");

      await store.getState().initialize();

      expect(store.getState().status).toBe("authenticated");
      expect(store.getState().connectedUser).not.toBeNull();
    });

    it("should only restore the session once, even if called concurrently", async () => {
      const { store, deps } = setup({ isUsersSlow: true });
      deps.tokenStorage.saveTokens("refresh", "access");
      const whoami = vi.spyOn(deps.authGateway, "whoami");

      const first = store.getState().initialize();
      expect(store.getState().status).toBe("checking");
      await Promise.all([first, store.getState().initialize()]);
      await store.getState().initialize();

      expect(whoami).toHaveBeenCalledOnce();
    });

    it("should keep the error and be unauthenticated if the API is down", async () => {
      const { store, deps } = setup({ isResponseError: true });
      deps.tokenStorage.saveTokens("refresh", "access");

      await store.getState().initialize();

      expect(store.getState().status).toBe("unauthenticated");
      expect(store.getState().sessionError).toBe("whoami error");
    });
  });

  it("login should authenticate", async () => {
    const { store } = setup();
    expect(await store.getState().login({ email: "a@b.c", password: "pwd" })).toBe(true);
    expect(store.getState().status).toBe("authenticated");
    expect(store.getState().loginError).toBeNull();
  });

  it("login errors should not leak into signup state", async () => {
    const { store } = setup({ isWrongPassword: true });
    await store.getState().login({ email: "a@b.c", password: "x" });
    expect(store.getState().loginError).toBe("invalid_credentials");
    expect(store.getState().signupError).toBeNull();
  });

  it("logout should clear the user and the tokens", async () => {
    const { store, deps } = setup();
    await store.getState().login({ email: "a@b.c", password: "pwd" });

    store.getState().logout();

    expect(store.getState().status).toBe("unauthenticated");
    expect(store.getState().connectedUser).toBeNull();
    expect(deps.tokenStorage.hasSession()).toBe(false);
  });

  it("should become unauthenticated when the session expires", async () => {
    const { store, deps } = setup();
    await store.getState().login({ email: "a@b.c", password: "pwd" });

    deps.sessionEvents.emitSessionExpired();

    expect(store.getState().status).toBe("unauthenticated");
    expect(store.getState().connectedUser).toBeNull();
  });

  describe("createWorkspace", () => {
    it("should add the created workspace to the connected user", async () => {
      const { store } = setup();
      await store.getState().login({ email: "a@b.c", password: "pwd" });
      const before = store.getState().connectedUser!.workspaces.length;

      expect(await store.getState().createWorkspace("  My kitchen  ")).toBe(true);

      const workspaces = store.getState().connectedUser!.workspaces;
      expect(workspaces).toHaveLength(before + 1);
      expect(workspaces.at(-1)!.name).toBe("My kitchen");
    });

    it("should reject a name outside the backend limits", async () => {
      const { store, deps } = setup();
      expect(await store.getState().createWorkspace("short")).toBe(false);
      expect(store.getState().createWorkspaceError).toMatch(/between 8 and 30/);
      expect(deps.workspaces.created).toHaveLength(0);
    });

    it("should expose the API error", async () => {
      const { store, deps } = setup();
      deps.workspaces.setIsResponseError(true);
      expect(await store.getState().createWorkspace("My kitchen")).toBe(false);
      expect(store.getState().createWorkspaceError).toBe("Could not create the workspace");
    });
  });
});
