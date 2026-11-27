import { TokenStorage } from "../domain/tokenStorage";

const ACCESS_TOKEN_KEY = "access_token";
const REFRESH_TOKEN_KEY = "refresh_token";

const isBrowser = () => typeof window !== "undefined";

export class LocalStorageTokenStorage implements TokenStorage {
  getAccessToken(): string | null {
    if (!isBrowser()) return null;
    return localStorage.getItem(ACCESS_TOKEN_KEY);
  }

  getRefreshToken(): string | null {
    if (!isBrowser()) return null;
    return localStorage.getItem(REFRESH_TOKEN_KEY);
  }

  saveTokens(refreshToken: string, accessToken: string): void {
    if (!isBrowser()) return;
    localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
    localStorage.setItem(REFRESH_TOKEN_KEY, refreshToken);
  }

  setAccessToken(accessToken: string): void {
    if (!isBrowser()) return;
    localStorage.setItem(ACCESS_TOKEN_KEY, accessToken);
  }

  removeTokens(): void {
    if (!isBrowser()) return;
    localStorage.removeItem(ACCESS_TOKEN_KEY);
    localStorage.removeItem(REFRESH_TOKEN_KEY);
  }

  // The refresh token is what keeps a session alive: the access token may be expired.
  hasSession(): boolean {
    return !!this.getRefreshToken();
  }

  /** Fires when another tab removes the session (logout or expiry). Returns an unsubscribe. */
  onRemovedInOtherTab(listener: () => void): () => void {
    if (!isBrowser()) return () => {};
    const handler = (event: StorageEvent) => {
      const cleared = event.key === null; // localStorage.clear()
      if (cleared || (event.key === REFRESH_TOKEN_KEY && event.newValue === null)) {
        listener();
      }
    };
    window.addEventListener("storage", handler);
    return () => window.removeEventListener("storage", handler);
  }
}
