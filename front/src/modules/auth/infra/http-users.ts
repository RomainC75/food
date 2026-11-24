import { AxiosInstance } from "axios";
import { Users } from "../domain/users";
import { axiosInstance } from "./http_handler";
import { ConnectedUserInfo, SignedUpUser, SigninUser, SignupUser } from "../domain/user";

export const ErrorUserUnverified = "Account not verified, please check your emails";

const signupUrl = "/id-access"

export class HttpUsers implements Users {
  constructor(
    private readonly httpHandler: AxiosInstance = axiosInstance,
  ) {}
  async signup(signupUser: SignupUser): Promise<SignedUpUser> {
    const response = await this.httpHandler.post<SignedUpUser>(
        `${signupUrl}/signup`,
        signupUser
    )
    return response.data
  }

  async login(newUser: SigninUser): Promise<ConnectedUserInfo> {
    const response = await this.httpHandler.post<ConnectedUserInfo>(
        `${signupUrl}/signup`,
        newUser
    )
    return response.data
  }

}
