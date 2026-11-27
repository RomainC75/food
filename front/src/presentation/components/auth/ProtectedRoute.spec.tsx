import { cleanup, render, screen, waitFor } from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import { act } from "react";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies, TestDependencies, TestDependenciesOptions } from "@/dependencies";
import ProtectedRoutesComp, { LOGIN_URL } from "./ProtectedRoutes";
import { ConnectedUserInfo } from "@/modules/auth/domain/user";

const replaceMock = vi.fn();
let pathname = "/";

vi.mock("next/navigation", () => ({
  useRouter: () => ({
    push: vi.fn(),
    replace: replaceMock,
    back: vi.fn(),
    refresh: vi.fn(),
    forward: vi.fn(),
    prefetch: vi.fn(),
  }),
  usePathname: () => pathname,
}));

describe("ProtectedRoute component", () => {
  beforeEach(() => {
    replaceMock.mockReset();
    pathname = "/";
  });
  afterEach(() => cleanup());

  it("should redirect to login if there is no session", async () => {
    Render();
    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith(LOGIN_URL));
    expect(screen.queryByText("AUTHORIZED")).toBeNull();
  });

  it("should keep the current page in the login redirect", async () => {
    pathname = "/new-workspace";
    Render();
    await waitFor(() =>
      expect(replaceMock).toHaveBeenCalledWith(`${LOGIN_URL}?next=%2Fnew-workspace`),
    );
  });

  it("should display children if already authenticated", async () => {
    Render({}, connectedUser);
    expect(await screen.findByText("AUTHORIZED")).toBeVisible();
    expect(replaceMock).not.toHaveBeenCalled();
  });

  it("should show a spinner while restoring the session, then the children", async () => {
    const deps = Render({ isUsersSlow: true }, undefined, (d) =>
      d.tokenStorage.saveTokens("refresh", "access"),
    );
    expect(screen.getByRole("status", { name: /loading/i })).toBeVisible();
    expect(await screen.findByText("AUTHORIZED")).toBeVisible();
    expect(replaceMock).not.toHaveBeenCalled();
    expect(deps.tokenStorage.getRefreshToken()).toBe("refresh");
  });

  it("should redirect to login and clear tokens if the session expired", async () => {
    const deps = Render({ isSessionExpired: true }, undefined, (d) =>
      d.tokenStorage.saveTokens("refresh", "access"),
    );
    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith(LOGIN_URL));
    expect(deps.tokenStorage.hasSession()).toBe(false);
  });

  it("should redirect to login when the session expires while on the page", async () => {
    const deps = Render({}, connectedUser);
    expect(await screen.findByText("AUTHORIZED")).toBeVisible();

    act(() => deps.sessionEvents.emitSessionExpired());

    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith(LOGIN_URL));
    expect(screen.queryByText("AUTHORIZED")).toBeNull();
  });
});

const connectedUser: ConnectedUserInfo = {
  id: "my-id",
  email: "fake@email.com",
  workspaces: [],
};

const Render = (
  opts: TestDependenciesOptions = {},
  user?: ConnectedUserInfo,
  setup?: (deps: TestDependencies) => void,
): TestDependencies => {
  const deps = createTestDependencies(opts);
  setup?.(deps);
  render(
    <Providers deps={deps} connectedUser={user}>
      <ProtectedRoutesComp>
        <div>AUTHORIZED</div>
      </ProtectedRoutesComp>
    </Providers>,
  );
  return deps;
};
