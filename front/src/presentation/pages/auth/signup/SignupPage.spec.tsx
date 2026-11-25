import { cleanup, render, screen } from "@testing-library/react";
import { afterEach, assert, describe, expect, it } from "vitest";
import userEvent from "@testing-library/user-event";
import SignupComponent from "./SignupPage";
import { vi } from "vitest";
import Providers from "@/presentation/state-manager/providers";
import { createTestDependencies, TestDependencies } from "@/dependencies";

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
  it("should show invalid email error as soon as the user enters somthing ", async () => {
    Render(false, false);
    const user = userEvent.setup();
    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "m");

    const res = await screen.findByText(/invalid email address/i);
    assert.exists(res);

    expect(screen.getByRole("button", { name: /SIGN UP/i })).toBeDisabled();
  });

  it("should show 'password must contains at least 8 characters'", async () => {
    Render(false,false);
    const user = userEvent.setup();

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "m");

    expect(
      await screen.findByText(/password must contain at least 8 characters/i),
    ).toBeVisible();

    expect(screen.getByRole("button", { name: /SIGN UP/i })).toBeDisabled();
  });

  it("should show 'Passwords do not match' is passwords are different", async () => {
    Render(false,false);
    const user = userEvent.setup();

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "fakePassword");

    const confirmPasswordInput = screen.getByLabelText("Confirm password");
    await user.type(confirmPasswordInput, "fakePasswo");

    expect(await screen.findByText(/Passwords do not match/i)).toBeVisible();

    expect(screen.getByRole("button", { name: /SIGN UP/i })).toBeDisabled();
  });

  it("should no signup if email is invalid then wait for response", async () => {
    Render(true, false);
    const user = userEvent.setup();

    const firstNameInput = screen.getByRole("textbox", {
      name: /first name/i,
    });
    await user.type(firstNameInput, "my-first-name");
    const lastNameInput = screen.getByRole("textbox", {
      name: /last name/i,
    });
    await user.type(lastNameInput, "my-last-name");

    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "fake@email.com");

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "my-password");

    const confirmPassword = screen.getByLabelText("Confirm password");
    await user.type(confirmPassword, "my-password");

    expect(screen.getByRole("button", { name: /SIGN UP/i })).toBeEnabled();
    await user.click(screen.getByRole("button", { name: /sign up/i }));
    
    expect(await screen.findByText(/Creating account/i)).toBeVisible();
  });

  it("should show a message if the signup returns an error", async ()=>{
        Render(false, true);
    const user = userEvent.setup();

    const firstNameInput = screen.getByRole("textbox", {
      name: /first name/i,
    });
    await user.type(firstNameInput, "my-first-name");
    const lastNameInput = screen.getByRole("textbox", {
      name: /last name/i,
    });
    await user.type(lastNameInput, "my-last-name");

    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "fake@email.com");

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "my-password");

    const confirmPassword = screen.getByLabelText("Confirm password");
    await user.type(confirmPassword, "my-password");

    expect(screen.getByRole("button", { name: /SIGN UP/i })).toBeEnabled();
    await user.click(screen.getByRole("button", { name: /sign up/i }));
    
    expect(await screen.findByText(/Error trying to signup, try again/i)).toBeVisible();
  })
    it("should go to / signedup", async ()=>{
        Render(false, false);
    const user = userEvent.setup();

    const firstNameInput = screen.getByRole("textbox", {
      name: /first name/i,
    });
    await user.type(firstNameInput, "my-first-name");
    const lastNameInput = screen.getByRole("textbox", {
      name: /last name/i,
    });
    await user.type(lastNameInput, "my-last-name");

    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "fake@email.com");

    const passwordInput = screen.getByLabelText("Password");
    await user.type(passwordInput, "my-password");

    const confirmPassword = screen.getByLabelText("Confirm password");
    await user.type(confirmPassword, "my-password");

    expect(screen.getByRole("button", { name: /SIGN UP/i })).toBeEnabled();
    await user.click(screen.getByRole("button", { name: /sign up/i }));

    expect(pushMock).toHaveBeenCalledWith("/");
  })
});



let deps: TestDependencies;

const Render = (isUsersSlow: boolean, isResponseError: boolean) => {
  deps = createTestDependencies({ isUsersSlow, isResponseError });

  render(
    <Providers deps={deps}>
      <SignupComponent />
    </Providers>,
  );
};
