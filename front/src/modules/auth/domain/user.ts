
export type SignupUser = {
    firstName: string,
    lastName: string,
    email: string,
    password: string
}

export type SigninUser =  {
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
}

export type FullConnectedUserInfo = ConnectedUserInfo & {
    token: string,
    refreshToken: string,
}