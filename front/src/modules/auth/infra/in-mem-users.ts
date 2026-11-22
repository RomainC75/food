import { Uuids } from "@/modules/shared/domain/uuid"
import { ConnectedUserInfo, SignedUpUser, UserConnection } from "../domain/user"
import { Users } from "../domain/users";

export class InMemUser implements Users{
    fakeToken = "FAKE_TOKEN";
    workdspaceId = ["WORKSPACE_1", "WORKSPACE_2"];

    constructor(private uuidGen: Uuids){}

    // ? for the moment, the id/token do not matter
    // ? don't let the NewUser like this
    signup(userConnection: UserConnection): Promise<SignedUpUser>{
        return Promise.resolve({
            id: this.uuidGen.now(),
            email: userConnection.email,
        })
    }

    login(newUser: UserConnection): Promise<ConnectedUserInfo>{
        return Promise.resolve({
            id: this.uuidGen.now(),
            email: newUser.email,
            workspaceIds: this.workdspaceId,
            token: this.fakeToken
        })
    }
}