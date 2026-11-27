export type WorkspaceResDto = {
  id: string;
  name: string;
};

export type WhoAmIResDto = {
  id: string;
  email: string;
  // Not sent by the backend yet (WhoAmIDto only has id + email)
  workspaces?: WorkspaceResDto[];
};

export type LoginResDto = {
  id: string;
  email: string;
  workspaces: WorkspaceResDto[] | null;
  access_token: string;
  refresh_token: string;
};

export type RefreshResDto = {
  access_token: string;
};

export type SignupResDto = {
  id: string;
  email: string;
};
