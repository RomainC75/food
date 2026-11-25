import { Uuids } from "@/modules/shared/domain/uuid";
import {
  FullConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "../domain/user";
import { Users } from "../domain/users";

export const ErrorUserUnverified = "Account not verified, please check your emails";
export const FAKE_TOKEN = "FAKE_TOKEN"

export class InMemUser implements Users {
  signedUpUser: SignupUser | null;
  fakeToken = FAKE_TOKEN;
  workdspaceId = ["WORKSPACE_1", "WORKSPACE_2"];
  isResponseError = false;
  isUserUnverified = false;
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

  login(newUser: SigninUser): Promise<FullConnectedUserInfo> {
    if (this.isUserUnverified){
      throw new Error(ErrorUserUnverified)
    }
    if (this.isResponseError) {
      throw new Error("login error");
    }
    return Promise.resolve({
      id: this.uuidGen.now(),
      email: newUser.email,
      workspaceIds: this.workdspaceId,
      accessToken: this.fakeToken,
      refreshToken: this.fakeToken
    });
  }

  setIsResponseError(isResponseError: boolean): InMemUser {
    this.isResponseError = isResponseError;
    return this;
  }
  setIsUserUnverified(isUserUnverified: boolean): InMemUser {
    this.isUserUnverified = isUserUnverified;
    return this;
  }
}
