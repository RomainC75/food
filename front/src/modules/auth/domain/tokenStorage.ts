export interface TokenStorage extends TokenStorageWriter, TokenStorageHandler {};
export interface TokenStorageHandler {
    getAccessToken(): string | null;
    getRefreshToken(): string | null;
    removeTokens(): void;
    setAccessToken(accessToken: string):void; 
}

export interface TokenStorageWriter{
    saveTokens(refreshToken: string, accessToken: string): void;
    isTokenInMemory(): boolean
}

