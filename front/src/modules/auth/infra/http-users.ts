import { AxiosInstance } from "axios";
import { Users } from "../domain/users";
import { axiosHttpHander } from "./http_handler";
import { FullConnectedUserInfo, SignedUpUser, SigninUser, SignupUser } from "../domain/user";
import { TokenStorage, TokenStorageWriter } from "../domain/tokenStorage";

export const ErrorUserUnverified = "Account not verified, please check your emails";

const signupUrl = "/id-access"

export class HttpUsers implements Users {
  httpHandler: AxiosInstance;
  constructor(
    private tokenStorage: TokenStorage,
  ) {
    this.httpHandler = axiosHttpHander(tokenStorage);
  }
  async signup(signupUser: SignupUser): Promise<SignedUpUser> {
    const response = await this.httpHandler.post<SignedUpUser>(
        `${signupUrl}/signup`,
        signupUser
    )
    return response.data
  }

  async login(newUser: SigninUser): Promise<FullConnectedUserInfo> {
    const response = await this.httpHandler.post<FullConnectedUserInfo>(
        `${signupUrl}/signup`,
        newUser
    )

    return response.data
  }

}
