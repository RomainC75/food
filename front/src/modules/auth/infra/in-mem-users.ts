import { Uuids } from "@/modules/shared/domain/uuid";
import {
  ConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "../domain/user";
import { Users } from "../domain/users";

export class InMemUser implements Users {
  signedUpUser: SignupUser | null;
  fakeToken = "FAKE_TOKEN";
  workdspaceId = ["WORKSPACE_1", "WORKSPACE_2"];
  isResponseError = false;
  constructor(
    private uuidGen: Uuids,
    private isUsersSlow: boolean,
  ) {
    this.signedUpUser = null;
  }

  // ? for the moment, the id/token do not matter
  // ? don't let the NewUser like this
  async signup(signupUser: SignupUser): Promise<SignedUpUser> {
    if (this.isResponseError) {
      throw new Error("signup error");
    }
    this.signedUpUser = signupUser;
    if (this.isUsersSlow) {
      await new Promise((resolve) => setTimeout(resolve, 10));
    }

    return Promise.resolve({
      id: this.uuidGen.now(),
      email: signupUser.email,
    });
  }

  login(newUser: SigninUser): Promise<ConnectedUserInfo> {
    if (this.isResponseError) {
      throw new Error("login error");
    }
    return Promise.resolve({
      id: this.uuidGen.now(),
      email: newUser.email,
      workspaceIds: this.workdspaceId,
      token: this.fakeToken,
    });
  }

  setIsResponseError(isResponseError: boolean): InMemUser {
    this.isResponseError = isResponseError;
    return this;
  }
}
