import axios, { AxiosInstance } from "axios";
import { TokenStorageHandler } from "../domain/tokenStorage";

export const axiosHttpHander = (tokenStorage: TokenStorageHandler): AxiosInstance => {

    const api = axios.create({
      baseURL: process.env.NEXT_PUBLIC_API_URL,
      headers: {
        "Content-Type": "application/json",
      },
    });

    // Request interceptor
    api.interceptors.request.use((config) => {
      const token = tokenStorage.getAccessToken();

      if (token) {
        config.headers.Authorization = `Bearer ${token}`;
      }

      return config;
    });

    // Response interceptor
    api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          tokenStorage.removeAccessToken();

          // Notify the frontend authentication state
          if (typeof window !== "undefined") {
            window.dispatchEvent(new Event("auth:unauthorized"));
          }
        }

        return Promise.reject(error);
      },
    );
    return api
} 



