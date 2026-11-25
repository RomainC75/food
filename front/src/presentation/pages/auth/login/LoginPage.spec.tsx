import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import LoginComponent from "./LoginPage";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies, TestDependencies } from "@/dependencies";
import { FAKE_TOKEN } from "@/modules/auth/infra/in-mem-users";


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

describe("SignupForm tests", () => {
  afterEach(() => cleanup());
  it("should show a message if the user is not verified", async () => {
    Render(false, false, true);
    const user = userEvent.setup();

    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "fake@email.com");

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "my-password");

    await user.click(screen.getByRole("button", { name: /login/i }));

    expect(await screen.findByText(/check your emails/i)).toBeVisible();
  });

    it("should go to / if login returns the token", async () => {
    const deps = Render(false, false, false);
    const user = userEvent.setup();

    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "fake@email.com");

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "my-password");
                 
    await user.click(screen.getByRole("button", { name: /login/i }));
    
    expect(pushMock).toHaveBeenCalledWith("/");

    // TODO test the token has been saved 
    expect(deps.tokenStorage.getAccessToken()).toEqual(FAKE_TOKEN)
  });
});

let deps: TestDependencies;

const Render = (
  isUsersSlow: boolean,
  isResponseError: boolean,
  isUserUnverified: boolean,
): TestDependencies => {
  deps = createTestDependencies({ isUsersSlow, isResponseError, isUserUnverified });

  render(
    <Providers deps={deps}>
      <LoginComponent />
    </Providers>,
  );
  return deps
};
