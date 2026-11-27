import { TokenStorage } from "../domain/tokenStorage";

export class InMemTokenStorage implements TokenStorage {
  accessToken: string | null = null;
  refreshToken: string | null = null;

  getAccessToken(): string | null {
    return this.accessToken;
  }

  getRefreshToken(): string | null {
    return this.refreshToken;
  }

  saveTokens(refreshToken: string, accessToken: string): void {
    this.refreshToken = refreshToken;
    this.accessToken = accessToken;
  }

  setAccessToken(accessToken: string): void {
    this.accessToken = accessToken;
  }

  removeTokens(): void {
    this.accessToken = null;
    this.refreshToken = null;
  }

  hasSession(): boolean {
    return !!this.refreshToken;
  }
}
