export interface TokenStorage extends TokenStorageWriter, TokenStorageHandler {};
export interface TokenStorageHandler {
    getAccessToken(): string | null;
    removeAccessToken(): void;
}

export interface TokenStorageWriter{
    saveAccessToken(token: string): void;
}

