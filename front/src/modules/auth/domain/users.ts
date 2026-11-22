import { ConnectedUserInfo, SignedUpUser, UserConnection } from "./user";



export interface Users{
    signup(userConnection: UserConnection): Promise<SignedUpUser>
    login(newUser: UserConnection): Promise<ConnectedUserInfo>
}