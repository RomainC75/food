import { describe, expect, it } from "vitest";
import { createTestDependencies } from "@/dependencies";
import { FAKE_TOKEN } from "../infra/in-mem-authGateway";
import { InvalidCredentialsError } from "../domain/errors";

describe("AuthSrv", () => {
  describe("login", () => {
    it("should save both tokens and return the user", async () => {
      const { authSrv, tokenStorage } = createTestDependencies();

      const user = await authSrv.login({ email: "a@b.c", password: "pwd" });

      expect(user.email).toBe("a@b.c");
      expect(tokenStorage.getAccessToken()).toBe(FAKE_TOKEN);
      expect(tokenStorage.getRefreshToken()).toBe(FAKE_TOKEN);
    });

    it("should not save anything on wrong credentials", async () => {
      const { authSrv, tokenStorage } = createTestDependencies({ isWrongPassword: true });

      await expect(authSrv.login({ email: "a@b.c", password: "x" })).rejects.toBeInstanceOf(
        InvalidCredentialsError,
      );
      expect(tokenStorage.hasSession()).toBe(false);
    });
  });

  describe("restoreSession", () => {
    it("should return null without calling the API when there is no session", async () => {
      const { authSrv, authGateway } = createTestDependencies();
      authGateway.setIsResponseError(true); // would throw if called

      await expect(authSrv.restoreSession()).resolves.toBeNull();
    });

    it("should return the user when the session is valid", async () => {
      const { authSrv, authGateway, tokenStorage } = createTestDependencies();
      tokenStorage.saveTokens("refresh", "access");
      authGateway.whoamiUser = { id: "id-1", email: "me@x.y", workspaces: [] };

      const user = await authSrv.restoreSession();

      expect(user).toMatchObject({ id: "id-1", email: "me@x.y" });
    });

    it("should clear the tokens and return null when the session expired", async () => {
      const { authSrv, tokenStorage } = createTestDependencies({ isSessionExpired: true });
      tokenStorage.saveTokens("refresh", "access");

      await expect(authSrv.restoreSession()).resolves.toBeNull();
      expect(tokenStorage.hasSession()).toBe(false);
    });

    it("should rethrow other errors and keep the tokens (e.g. API down)", async () => {
      const { authSrv, tokenStorage } = createTestDependencies({ isResponseError: true });
      tokenStorage.saveTokens("refresh", "access");

      await expect(authSrv.restoreSession()).rejects.toThrow("whoami error");
      expect(tokenStorage.hasSession()).toBe(true);
    });
  });

  it("logout should clear the tokens", async () => {
    const { authSrv, tokenStorage } = createTestDependencies();
    tokenStorage.saveTokens("refresh", "access");

    authSrv.logout();

    expect(tokenStorage.hasSession()).toBe(false);
  });
});
