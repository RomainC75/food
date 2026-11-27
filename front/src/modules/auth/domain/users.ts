import { WhoAmIResDto } from "@/modules/dtos/res/auth-res";
import {
  FullConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "./user";

export interface Users {
  signup(userConnection: SignupUser): Promise<SignedUpUser>;
  login(newUser: SigninUser): Promise<FullConnectedUserInfo>;
  refresh():Promise<void>
  whoami(): Promise<WhoAmIResDto>
  //   logout(): Promise<void>;
  //   getCurrentUser(): Promise<User | null>;
}
