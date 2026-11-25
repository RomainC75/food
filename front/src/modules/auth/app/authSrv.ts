import { TokenStorageWriter } from "../domain/tokenStorage";
import { ConnectedUserInfo, SignedUpUser, SigninUser, SignupUser } from "../domain/user";
import { Users } from "../domain/users";

export class AuthSrv {
    constructor(
        private users: Users,
        private tokenStorage: TokenStorageWriter
    ){ }

    signup(userConnection: SignupUser): Promise<SignedUpUser>{
        return this.users.signup(userConnection)
    }

    async login(userConnection: SigninUser): Promise<ConnectedUserInfo>{
        const fullUserInfo = await this.users.login(userConnection)
        this.tokenStorage.saveAccessToken(fullUserInfo.refreshToken)
        return {
            id: fullUserInfo.id,
            email: fullUserInfo.email,
            workspaceIds: fullUserInfo.workspaceIds,
        }
    }
}