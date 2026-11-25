import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, describe, expect, it, vi } from "vitest";
import userEvent from "@testing-library/user-event";
import LoginComponent from "./LoginForm";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies, Dependencies, TestDependencies } from "@/dependencies";


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
    const dependencies = Render(false, false, false);
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

  });
});

let deps: TestDependencies;

const Render = (
  isUsersSlow: boolean,
  isResponseError: boolean,
  isUserUnverified: boolean,
): Dependencies => {
  deps = createTestDependencies({ isUsersSlow, isResponseError, isUserUnverified });

  render(
    <Providers deps={deps}>
      <LoginComponent />
    </Providers>,
  );
  return deps
};
