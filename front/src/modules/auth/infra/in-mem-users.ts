import { Uuids } from "@/modules/shared/domain/uuid";
import {
  FullConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
  Workspace,
} from "../domain/user";
import { Users } from "../domain/users";
import { WhoAmIResDto } from "@/modules/dtos/res/auth-res";

export const ErrorUserUnverified =
  "Account not verified, please check your emails";
export const FAKE_TOKEN = "FAKE_TOKEN";

export class InMemUser implements Users {
  refreshEmail = "";
  refreshId = "";
  signedUpUser: SignupUser | null;
  fakeToken = FAKE_TOKEN;
  workdspaces: Workspace[] = [{id: "1", name: "WORKSPACE_1"}, {id: "2", name: "WORKSPACE_2"}];
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
    if (this.isUserUnverified) {
      throw new Error(ErrorUserUnverified);
    }
    if (this.isResponseError) {
      throw new Error("login error");
    }
    return Promise.resolve({
      id: this.uuidGen.now(),
      email: newUser.email,
      workspaces: this.workdspaces,
      accessToken: this.fakeToken,
      refreshToken: this.fakeToken,
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

  refresh(): Promise<void> {
    return Promise.resolve();
  }

  whoami(): Promise<WhoAmIResDto> {
    return Promise.resolve({
      id: this.refreshId,
      email: this.refreshEmail,
      workspaces: this.workdspaces,
    });
  }
}
