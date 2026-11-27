export class InvalidCredentialsError extends Error {
  constructor(message = "Wrong email or password") {
    super(message);
    this.name = "InvalidCredentialsError";
  }
}

export class UserUnverifiedError extends Error {
  constructor(message = "Account not verified, please check your emails") {
    super(message);
    this.name = "UserUnverifiedError";
  }
}

// The session can't be restored or refreshed anymore: the user has to log in again.
export class SessionExpiredError extends Error {
  constructor(message = "Session expired") {
    super(message);
    this.name = "SessionExpiredError";
  }
}
