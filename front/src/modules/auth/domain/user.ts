export type SignupUser = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
};

export type SigninUser = {
  email: string;
  password: string;
};

export type SignedUpUser = {
  id: string;
  email: string;
};

export type ConnectedUserInfo = {
  id: string;
  email: string;
  workspaces: Workspace[];
};

export type Workspace = {
  id: string;
  name: string;
};

export type AuthTokens = {
  accessToken: string;
  refreshToken: string;
};

export type LoginResult = {
  user: ConnectedUserInfo;
  tokens: AuthTokens;
};
