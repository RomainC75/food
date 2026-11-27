import { Uuids } from "@/modules/shared/domain/uuid";
import { Workspaces } from "../domain/workspaces";
import { Workspace } from "../domain/user";

export class InMemWorkspaces implements Workspaces {
  created: Workspace[] = [];
  isResponseError = false;

  constructor(private uuidGen: Uuids) {}

  async create(name: string): Promise<Workspace> {
    if (this.isResponseError) throw new Error("Could not create the workspace");
    const workspace = { id: this.uuidGen.now(), name };
    this.created.push(workspace);
    return workspace;
  }

  setIsResponseError(isResponseError: boolean): InMemWorkspaces {
    this.isResponseError = isResponseError;
    return this;
  }
}
