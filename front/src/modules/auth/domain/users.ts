import {
  ConnectedUserInfo,
  FullConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "./user";

export interface Users {
  signup(userConnection: SignupUser): Promise<SignedUpUser>;
  login(newUser: SigninUser): Promise<FullConnectedUserInfo>;

  //   logout(): Promise<void>;
  //   getCurrentUser(): Promise<User | null>;
  //   refresh(): Promise<void>;
}
