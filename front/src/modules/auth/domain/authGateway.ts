import {
  ConnectedUserInfo,
  LoginResult,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "./user";

export interface AuthGateway {
  signup(user: SignupUser): Promise<SignedUpUser>;
  /** @throws InvalidCredentialsError | UserUnverifiedError */
  login(credentials: SigninUser): Promise<LoginResult>;
  /** @throws SessionExpiredError when the session can't be refreshed */
  whoami(): Promise<ConnectedUserInfo>;
}
