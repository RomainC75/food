import axios, { AxiosInstance } from "axios";
import { TokenStorageHandler } from "../domain/tokenStorage";

declare module "axios" {
  export interface AxiosRequestConfig {
    skipAuth?: boolean;
  }
}

export const axiosHttpHander = (
  tokenStorage: TokenStorageHandler,
): AxiosInstance => {
  const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
    headers: {
      "Content-Type": "application/json",
    },
  });

  // Request interceptor
  api.interceptors.request.use((config) => {
    if (config.skipAuth) {
      return config;
    }
    const accesToken = tokenStorage.getAccessToken();

    if (accesToken) {
      config.headers.Authorization = `Bearer ${accesToken}`;
    }

    return config;
  });

  // Response interceptor
  api.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response?.status === 401) {
        tokenStorage.removeTokens();

        // Notify the frontend authentication state
        if (typeof window !== "undefined") {
          window.dispatchEvent(new Event("auth:unauthorized"));
        }
      }

      return Promise.reject(error);
    },
  );
  return api;
};
