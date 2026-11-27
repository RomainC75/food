import { Uuids } from "@/modules/shared/domain/uuid";
import {
  ConnectedUserInfo,
  LoginResult,
  SignedUpUser,
  SigninUser,
  SignupUser,
  Workspace,
} from "../domain/user";
import { AuthGateway } from "../domain/authGateway";
import {
  InvalidCredentialsError,
  SessionExpiredError,
  UserUnverifiedError,
} from "../domain/errors";

export const FAKE_REFRESH_TOKEN = "FAKE_REFRESH_TOKEN ";
export const FAKE_ACCESS_TOKEN = "FAKE_ACCESS_TOKEN";

export class InMemAuthGateway implements AuthGateway {
  signedUpUser: SignupUser | null = null;
  whoamiUser: ConnectedUserInfo = { id: "", email: "", workspaces: [] };
  workspaces: Workspace[] = [
    { id: "1", name: "WORKSPACE_1" },
    { id: "2", name: "WORKSPACE_2" },
  ];
  isResponseError = false;
  isUserUnverified = false;
  isWrongPassword = false;
  isSessionExpired = false;

  constructor(
    private uuidGen: Uuids,
    private isSlow: boolean,
  ) {}

  async signup(signupUser: SignupUser): Promise<SignedUpUser> {
    if (this.isResponseError) throw new Error("signup error");
    this.signedUpUser = signupUser;
    await this.wait();
    return { id: this.uuidGen.now(), email: signupUser.email };
  }

  async login(credentials: SigninUser): Promise<LoginResult> {
    if (this.isUserUnverified) throw new UserUnverifiedError();
    if (this.isWrongPassword) throw new InvalidCredentialsError();
    if (this.isResponseError) throw new Error("login error");
    await this.wait();
    return {
      user: {
        id: this.uuidGen.now(),
        email: credentials.email,
        workspaces: this.workspaces,
      },
      tokens: { accessToken: FAKE_ACCESS_TOKEN, refreshToken: FAKE_REFRESH_TOKEN },
    };
  }

  async whoami(): Promise<ConnectedUserInfo> {
    if (this.isSessionExpired) throw new SessionExpiredError();
    if (this.isResponseError) throw new Error("whoami error");
    await this.wait();
    return { ...this.whoamiUser, workspaces: this.workspaces };
  }

  setIsResponseError(isResponseError: boolean): InMemAuthGateway {
    this.isResponseError = isResponseError;
    return this;
  }

  setIsUserUnverified(isUserUnverified: boolean): InMemAuthGateway {
    this.isUserUnverified = isUserUnverified;
    return this;
  }

  setIsWrongPassword(isWrongPassword: boolean): InMemAuthGateway {
    this.isWrongPassword = isWrongPassword;
    return this;
  }

  setIsSessionExpired(isSessionExpired: boolean): InMemAuthGateway {
    this.isSessionExpired = isSessionExpired;
    return this;
  }

  private async wait() {
    if (this.isSlow) await new Promise((resolve) => setTimeout(resolve, 10));
  }
}
