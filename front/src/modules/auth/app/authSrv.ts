import { ConnectedUserInfo, SignedUpUser, UserConnection } from "../domain/user";
import { Users } from "../domain/users";

export class AuthSrv {
    constructor(private users: Users){ }

    signup(userConnection: UserConnection): Promise<SignedUpUser>{
        return this.users.signup(userConnection)
    }

    login(userConnection: UserConnection): Promise<ConnectedUserInfo>{
        return this.users.login(userConnection)
    }
}