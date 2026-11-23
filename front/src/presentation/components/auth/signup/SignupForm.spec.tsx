import { cleanup, render, screen } from "@testing-library/react";
import { beforeEach, afterEach, assert, describe, expect, it } from "vitest";
import userEvent from "@testing-library/user-event";
import { Providers } from "@/presentation/state-manager/providers";
import SignupComponent from "./SignupForm";

describe("SignupForm tests", () => {
    beforeEach(()=>{Render()});
    afterEach(() => cleanup());
  it("should show invalid email error as soon as the user enters somthing ", async () => {
    const user = userEvent.setup();
    const emailInput = screen.getByRole("textbox", {
      name: /email/i,
    });
    await user.type(emailInput, "m");

    const res = await screen.findByText(/invalid email address/i);
    assert.exists(res);
  });
  
  it("should show 'password must contains at least 8 characters'", async () => {
  
      const user = userEvent.setup();
      const passwordInput = screen.getByLabelText("Password");
      
      await user.type(passwordInput, "m");
  
      expect(
      await screen.findByText(/password must contain at least 8 characters/i)
    ).toBeVisible();
    });
});


// it("should no signup if email is invalid", async ()=>{
//     Render();

//     const user = userEvent.setup();
//     const firstNameInput = screen.getByRole("textbox", {
//         name: /first name/i,
//     });
//     await user.type(firstNameInput, "my-first-name");
//     const lastNameInput = screen.getByRole("textbox", {
//         name: /last name/i,
//     });
//     await user.type(lastNameInput, "my-last-name");

//     const passwordInput = screen.getByLabelText("Password");
//     await user.type(passwordInput, "my-password");

//     const confirmPassword = screen.getByLabelText("Confirm password");
//     await user.type(confirmPassword, "my-password");

//     await user.click(
//     screen.getByRole("button", { name: /sign up/i })
//     );
// })

// })

const Render = () => {
  render(
    <Providers initPlainIngredients={[]}>
      <SignupComponent />
    </Providers>,
  );
};
