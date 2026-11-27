import { cleanup, render, screen, waitFor } from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies } from "@/dependencies";
import GuestRoutesComp from "./GuestRoutes";
import { ConnectedUserInfo } from "@/modules/auth/domain/user";

const replaceMock = vi.fn();

vi.mock("next/navigation", () => ({
  useRouter: () => ({
    push: vi.fn(),
    replace: replaceMock,
    back: vi.fn(),
    refresh: vi.fn(),
    forward: vi.fn(),
    prefetch: vi.fn(),
  }),
}));

describe("GuestRoutes component", () => {
  beforeEach(() => {
    replaceMock.mockReset();
    window.history.replaceState(null, "", "/auth/login");
  });
  afterEach(() => cleanup());

  it("should display children if there is no session", async () => {
    Render();
    expect(await screen.findByText("LOGIN FORM")).toBeVisible();
    expect(replaceMock).not.toHaveBeenCalled();
  });

  it("should send an authenticated user to /", async () => {
    Render(connectedUser);
    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith("/"));
    expect(screen.queryByText("LOGIN FORM")).toBeNull();
  });

  it("should send an authenticated user to the 'next' page", async () => {
    window.history.replaceState(null, "", "/auth/login?next=%2Fnew-workspace");
    Render(connectedUser);
    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith("/new-workspace"));
  });

  it("should ignore an external 'next' url", async () => {
    window.history.replaceState(null, "", "/auth/login?next=%2F%2Fevil.com");
    Render(connectedUser);
    await waitFor(() => expect(replaceMock).toHaveBeenCalledWith("/"));
  });
});

const connectedUser: ConnectedUserInfo = {
  id: "my-id",
  email: "fake@email.com",
  workspaces: [],
};

const Render = (user?: ConnectedUserInfo) => {
  render(
    <Providers deps={createTestDependencies()} connectedUser={user}>
      <GuestRoutesComp>
        <div>LOGIN FORM</div>
      </GuestRoutesComp>
    </Providers>,
  );
};
