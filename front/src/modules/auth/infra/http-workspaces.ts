import { AxiosInstance } from "axios";
import { Workspaces } from "../domain/workspaces";
import { Workspace } from "../domain/user";
import { WorkspaceResDto } from "@/modules/dtos/res/auth-res";
import { httpErrorMessage } from "./http-client";

const WORKSPACE_URL = "/auth/workspace";

export class HttpWorkspaces implements Workspaces {
  constructor(private http: AxiosInstance) {}

  async create(name: string): Promise<Workspace> {
    try {
      const { data } = await this.http.post<WorkspaceResDto>(WORKSPACE_URL, { name });
      return { id: data.id, name: data.name };
    } catch (error) {
      throw new Error(httpErrorMessage(error, "Could not create the workspace"));
    }
  }
}
