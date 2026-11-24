import {
  ConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "./user";

export interface Users {
  signup(userConnection: SignupUser): Promise<SignedUpUser>;
  login(newUser: SigninUser): Promise<ConnectedUserInfo>;

  //   logout(): Promise<void>;
  //   getCurrentUser(): Promise<User | null>;
  //   refresh(): Promise<void>;
}
