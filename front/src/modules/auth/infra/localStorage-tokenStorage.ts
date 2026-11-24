import { TokenStorage } from "../domain/tokenStorage";
const ACCESS_TOKEN_KEY = "access_token";

export class LocalStorageTokenStorage implements TokenStorage{
    getAccessToken(): string | null{
        if (typeof window === "undefined") return null;
        return localStorage.getItem(ACCESS_TOKEN_KEY);
    }
    saveAccessToken(token: string): void{
        if (typeof window === "undefined") return;
        localStorage.setItem(ACCESS_TOKEN_KEY, token);
    }
    removeAccessToken(): void{
        if (typeof window === "undefined") return;
        localStorage.removeItem(ACCESS_TOKEN_KEY);
    }
}