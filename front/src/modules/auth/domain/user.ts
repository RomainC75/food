
export type UserConnection =  {
    email: string,
    password: string,
}

export type SignedUpUser = {
    id: string,
    email: string,
}

export type ConnectedUserInfo = {
    id: string, 
    workspaceIds: string[], 
    email: string, 
    token: string
}