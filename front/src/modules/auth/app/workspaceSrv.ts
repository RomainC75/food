import { Workspace } from "../domain/user";
import { Workspaces } from "../domain/workspaces";

// Must match the backend rule: validate:"required,gte=8,lte=30"
export const WORKSPACE_NAME_MIN = 8;
export const WORKSPACE_NAME_MAX = 30;

export class WorkspaceSrv {
  constructor(private workspaces: Workspaces) {}

  create(name: string): Promise<Workspace> {
    const trimmed = name.trim();
    if (trimmed.length < WORKSPACE_NAME_MIN || trimmed.length > WORKSPACE_NAME_MAX) {
      return Promise.reject(
        new Error(`Workspace name must be between ${WORKSPACE_NAME_MIN} and ${WORKSPACE_NAME_MAX} characters`),
      );
    }
    return this.workspaces.create(trimmed);
  }
}
