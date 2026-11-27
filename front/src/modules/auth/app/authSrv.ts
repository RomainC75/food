import { AuthGateway } from "../domain/authGateway";
import { SessionExpiredError } from "../domain/errors";
import { SessionEvents } from "../domain/sessionEvents";
import { TokenStorageWriter } from "../domain/tokenStorage";
import {
  ConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "../domain/user";

export class AuthSrv {
  constructor(
    private gateway: AuthGateway,
    private tokenStorage: TokenStorageWriter,
    private sessionEvents: SessionEvents,
  ) {}

  signup(user: SignupUser): Promise<SignedUpUser> {
    return this.gateway.signup(user);
  }

  async login(credentials: SigninUser): Promise<ConnectedUserInfo> {
    const { user, tokens } = await this.gateway.login(credentials);
    this.tokenStorage.saveTokens(tokens.refreshToken, tokens.accessToken);
    return user;
  }

  /**
   * Restores the session from stored tokens.
   * Returns null when there is no usable session (no tokens, or refresh rejected).
   * Other errors (network, 5xx) are rethrown and the tokens are kept.
   */
  async restoreSession(): Promise<ConnectedUserInfo | null> {
    if (!this.tokenStorage.hasSession()) return null;

    try {
      // An expired access token is refreshed transparently by the HTTP client
      return await this.gateway.whoami();
    } catch (error) {
      if (error instanceof SessionExpiredError) {
        this.tokenStorage.removeTokens();
        return null;
      }
      throw error;
    }
  }

  logout(): void {
    this.tokenStorage.removeTokens();
  }

  /** Session lost outside of a user action (refresh rejected, logout in another tab). */
  onSessionExpired(listener: () => void): () => void {
    return this.sessionEvents.onSessionExpired(listener);
  }
}
