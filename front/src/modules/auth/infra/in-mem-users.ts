import { Uuids } from "@/modules/shared/domain/uuid"
import { ConnectedUserInfo, SignedUpUser, SigninUser, SignupUser} from "../domain/user"
import { Users } from "../domain/users";

export class InMemUser implements Users{
    signedUpUser: SignupUser | null;
    fakeToken = "FAKE_TOKEN";
    workdspaceId = ["WORKSPACE_1", "WORKSPACE_2"];
    shouldWait = false
    constructor(private uuidGen: Uuids){
        this.signedUpUser = null
    }

    // ? for the moment, the id/token do not matter
    // ? don't let the NewUser like this
    async signup(signupUser: SignupUser): Promise<SignedUpUser>{
        this.signedUpUser = signupUser;
        if (this.shouldWait){
            await new Promise(resolve => setTimeout(resolve, 2000))
        }
        return Promise.resolve({
            id: this.uuidGen.now(),
            email: signupUser.email,
        })
    }

    login(newUser: SigninUser): Promise<ConnectedUserInfo>{
        return Promise.resolve({
            id: this.uuidGen.now(),
            email: newUser.email,
            workspaceIds: this.workdspaceId,
            token: this.fakeToken
        })
    }

    setShouldWait(shouldWait: boolean){
        this.shouldWait = shouldWait
    }
}