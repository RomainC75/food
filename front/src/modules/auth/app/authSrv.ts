import { ConnectedUserInfo, SignedUpUser, SigninUser, SignupUser } from "../domain/user";
import { Users } from "../domain/users";

export class AuthSrv {
    constructor(private users: Users){ }

    signup(userConnection: SignupUser): Promise<SignedUpUser>{
        return this.users.signup(userConnection)
    }

    login(userConnection: SigninUser): Promise<ConnectedUserInfo>{
        // TODO connect
        return this.users.login(userConnection)
    }
}