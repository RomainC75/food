import axios, {
  AxiosError,
  AxiosInstance,
  CreateAxiosDefaults,
  InternalAxiosRequestConfig,
} from "axios";
import { TokenStorageHandler } from "../domain/tokenStorage";
import { SessionExpiredError } from "../domain/errors";
import { RefreshResDto } from "@/modules/dtos/res/auth-res";

declare module "axios" {
  export interface AxiosRequestConfig {
    /** Don't attach the access token and don't try to refresh on 401 (login, signup, refresh). */
    skipAuth?: boolean;
    /** Internal: the request has already been replayed after a refresh. */
    _retry?: boolean;
  }
}

export const REFRESH_URL = "/auth/refresh";

type HttpClientOptions = {
  tokenStorage: TokenStorageHandler;
  /** Called once the session is definitely gone (refresh failed or replay still 401). */
  onSessionExpired: () => void;
  axiosConfig?: CreateAxiosDefaults;
};

export const createHttpClient = ({
  tokenStorage,
  onSessionExpired,
  axiosConfig,
}: HttpClientOptions): AxiosInstance => {
  const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
    headers: { "Content-Type": "application/json" },
    ...axiosConfig,
  });

  api.interceptors.request.use((config) => {
    if (config.skipAuth) return config;

    const accessToken = tokenStorage.getAccessToken();
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }
    return config;
  });

  // Single-flight: N requests failing with 401 at the same time share one refresh call.
  let refreshing: Promise<string> | null = null;

  const refreshAccessToken = async (): Promise<string> => {
    const refreshToken = tokenStorage.getRefreshToken();
    if (!refreshToken) throw new SessionExpiredError();

    // The backend answers 400 on an invalid refresh token: any failure ends the session.
    const { data } = await api.post<RefreshResDto>(
      REFRESH_URL,
      { refresh_token: refreshToken },
      { skipAuth: true },
    );
    tokenStorage.setAccessToken(data.access_token);
    return data.access_token;
  };

  const endSession = (): SessionExpiredError => {
    tokenStorage.removeTokens();
    onSessionExpired();
    return new SessionExpiredError();
  };

  api.interceptors.response.use(undefined, async (error: unknown) => {
    if (!axios.isAxiosError(error) || error.response?.status !== 401) throw error;

    const original = error.config as InternalAxiosRequestConfig | undefined;
    if (!original || original.skipAuth) throw error;

    // Already refreshed once and still 401: the session is not usable anymore.
    if (original._retry) throw endSession();

    let accessToken: string;
    try {
      refreshing ??= refreshAccessToken().finally(() => {
        refreshing = null;
      });
      accessToken = await refreshing;
    } catch {
      throw endSession();
    }

    original._retry = true;
    original.headers.Authorization = `Bearer ${accessToken}`;
    return api(original);
  });

  return api;
};

/** Text sent by the Go backend (http.Error writes plain text), or a fallback. */
export const httpErrorMessage = (error: unknown, fallback: string): string => {
  if (error instanceof AxiosError) {
    const data = error.response?.data;
    if (typeof data === "string" && data.trim()) return data.trim();
  }
  return error instanceof Error && error.message ? error.message : fallback;
};
