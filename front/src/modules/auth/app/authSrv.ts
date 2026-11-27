import { WhoAmIResDto } from "@/modules/dtos/res/auth-res";
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
        console.log("--> userSrv : login()", fullUserInfo)
        this.tokenStorage.saveTokens(fullUserInfo.refreshToken, fullUserInfo.accessToken)
        return {
            id: fullUserInfo.id,
            email: fullUserInfo.email,
            workspaces: fullUserInfo.workspaces,
        }
    }

    isTokenInMemory(): boolean {
        return this.tokenStorage.isTokenInMemory()
    }

    async refresh(): Promise<void>{
        console.log("--> refresh")
        await this.users.refresh()
    }

    async whoami(): Promise<WhoAmIResDto>{
        console.log("-> whoami")
        return this.users.whoami()
    }


}