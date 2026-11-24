import { TokenStorage } from "../domain/tokenStorage";

export class InMemTokenStorage implements TokenStorage{
    token: string | null = null;
    getAccessToken(): string | null{
        return this.token;
    }
    saveAccessToken(token: string): void{
        this.token = token;
    }
    removeAccessToken(): void{
        this.token = null;
    }
}