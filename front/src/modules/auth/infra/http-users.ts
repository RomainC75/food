import { AxiosInstance } from "axios";
import { Users } from "../domain/users";
import { axiosHttpHander } from "./http_handler";
import { FullConnectedUserInfo, SignedUpUser, SigninUser, SignupUser } from "../domain/user";
import { TokenStorage } from "../domain/tokenStorage";
import { LoginResDto, RefreshResDto, WhoAmIResDto } from "@/modules/dtos/res/auth-res";

export const ErrorUserUnverified = "Account not verified, please check your emails";

const signupUrl = "/auth"

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
        signupUser,
        {skipAuth: true}
    )
    return response.data
  }

  async login(newUser: SigninUser): Promise<FullConnectedUserInfo> {
    const {data} = await this.httpHandler.post<LoginResDto>(
        `${signupUrl}/login`,
        newUser,
        {skipAuth: true}
    )
    return {
      ...data,
      refreshToken: data.refresh_token,
      accessToken: data.access_token,
    }
  }

  async refresh():Promise<void>{
    const refreshToken = this.tokenStorage.getRefreshToken();
    const {data} = await this.httpHandler.post<RefreshResDto>(
      `${signupUrl}/refresh`,
      {
        refresh_token: refreshToken,
      },
      {skipAuth: true}
    );
    this.tokenStorage.setAccessToken(data.access_token)
  }

  async whoami(): Promise<WhoAmIResDto>{
    const {data} = await this.httpHandler.get<WhoAmIResDto>(
      `${signupUrl}/whoami`
    );
    return data;
  }
}
