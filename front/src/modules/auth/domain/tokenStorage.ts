export interface TokenStorage {
    getAccessToken(): string | null;
    saveAccessToken(token: string): void;
    removeAccessToken(): void;
}