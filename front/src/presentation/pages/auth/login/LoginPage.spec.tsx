import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import LoginComponent from "./LoginPage";
import Providers from "@/presentation/state-manager/providers";
import {
  createTestDependencies,
  TestDependencies,
  TestDependenciesOptions,
} from "@/dependencies";
import { FAKE_ACCESS_TOKEN, FAKE_REFRESH_TOKEN } from "@/modules/auth/infra/in-mem-authGateway";

const pushMock = vi.fn();

vi.mock("next/navigation", () => ({
  useRouter: () => ({
    push: pushMock,
    replace: vi.fn(),
    back: vi.fn(),
    refresh: vi.fn(),
    forward: vi.fn(),
    prefetch: vi.fn(),
  }),
}));

describe("LoginPage tests", () => {
  beforeEach(() => {
    pushMock.mockReset();
    window.history.replaceState(null, "", "/auth/login");
  });
  afterEach(() => cleanup());

  it("should show a message if the user is not verified", async () => {
    Render({ isUserUnverified: true });
    await fillAndSubmit();
    expect(await screen.findByText(/check your emails/i)).toBeVisible();
    expect(pushMock).not.toHaveBeenCalled();
  });

  it("should show a message if the email or password is wrong", async () => {
    Render({ isWrongPassword: true });
    await fillAndSubmit();
    expect(await screen.findByText(/wrong email or password/i)).toBeVisible();
  });

  it("should show a generic message on any other error", async () => {
    Render({ isResponseError: true });
    await fillAndSubmit();
    expect(await screen.findByText(/error trying to login, try again/i)).toBeVisible();
  });

  it("should save the tokens and go to / on success", async () => {
    const deps = Render();
    await fillAndSubmit();

    expect(pushMock).toHaveBeenCalledWith("/");
    expect(deps.tokenStorage.getAccessToken()).toEqual(FAKE_ACCESS_TOKEN);
    expect(deps.tokenStorage.getRefreshToken()).toEqual(FAKE_REFRESH_TOKEN);
  });

  it("should go back to the 'next' page on success", async () => {
    window.history.replaceState(null, "", "/auth/login?next=%2Fnew-workspace");
    Render();
    await fillAndSubmit();
    expect(pushMock).toHaveBeenCalledWith("/new-workspace");
  });
});

const fillAndSubmit = async () => {
  const user = userEvent.setup();
  await user.type(screen.getByRole("textbox", { name: /email/i }), "fake@email.com");
  await user.type(screen.getByLabelText("Password"), "my-password");
  await user.click(screen.getByRole("button", { name: /login/i }));
};

const Render = (opts: TestDependenciesOptions = {}): TestDependencies => {
  const deps = createTestDependencies(opts);
  render(
    <Providers deps={deps}>
      <LoginComponent />
    </Providers>,
  );
  return deps;
};
