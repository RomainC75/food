import { Workspace } from "@/modules/auth/domain/user";

export type WhoAmIResDto = {
  id: string;
  email: string;
  workspaces: Workspace[];
}

export type LoginResDto = WhoAmIResDto & {
  access_token: string;
  refresh_token: string;
};

export type RefreshResDto = {
  access_token: string;
}
