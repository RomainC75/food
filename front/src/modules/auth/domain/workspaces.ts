import { Workspace } from "./user";

export interface Workspaces {
  create(name: string): Promise<Workspace>;
}
