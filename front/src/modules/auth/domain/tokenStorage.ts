export interface TokenStorage extends TokenStorageWriter, TokenStorageHandler {}

// Used by the HTTP client
export interface TokenStorageHandler {
  getAccessToken(): string | null;
  getRefreshToken(): string | null;
  setAccessToken(accessToken: string): void;
  removeTokens(): void;
}

// Used by AuthSrv
export interface TokenStorageWriter {
  saveTokens(refreshToken: string, accessToken: string): void;
  hasSession(): boolean;
  removeTokens(): void;
}
