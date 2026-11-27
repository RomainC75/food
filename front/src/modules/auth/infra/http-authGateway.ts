import { AxiosError, AxiosInstance } from "axios";
import { AuthGateway } from "../domain/authGateway";
import {
  ConnectedUserInfo,
  LoginResult,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "../domain/user";
import { InvalidCredentialsError, UserUnverifiedError } from "../domain/errors";
import {
  LoginResDto,
  SignupResDto,
  WhoAmIResDto,
} from "@/modules/dtos/res/auth-res";

const AUTH_URL = "/auth";

export class HttpAuthGateway implements AuthGateway {
  constructor(private http: AxiosInstance) {}

  async signup(user: SignupUser): Promise<SignedUpUser> {
    const { data } = await this.http.post<SignupResDto>(
      `${AUTH_URL}/signup`,
      user,
      { skipAuth: true },
    );
    return { id: data.id, email: data.email };
  }

  async login(credentials: SigninUser): Promise<LoginResult> {
    try {
      const { data } = await this.http.post<LoginResDto>(
        `${AUTH_URL}/login`,
        credentials,
        { skipAuth: true },
      );
      return {
        user: {
          id: data.id,
          email: data.email,
          workspaces: data.workspaces ?? [],
        },
        tokens: {
          accessToken: data.access_token,
          refreshToken: data.refresh_token,
        },
      };
    } catch (error) {
      throw toLoginError(error);
    }
  }

  async whoami(): Promise<ConnectedUserInfo> {
    // A 401 here is handled by the HTTP client (refresh + replay, or SessionExpiredError)
    const { data } = await this.http.get<WhoAmIResDto>(`${AUTH_URL}/whoami`);
    return {
      id: data.id,
      email: data.email,
      workspaces: data.workspaces ?? [],
    };
  }
}

const toLoginError = (error: unknown): unknown => {
  if (!(error instanceof AxiosError)) return error;
  switch (error.response?.status) {
    case 403:
      return new UserUnverifiedError();
    case 400:
    case 401:
      return new InvalidCredentialsError();
    default:
      return error;
  }
};
