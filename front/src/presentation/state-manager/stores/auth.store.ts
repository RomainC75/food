import { createStore } from "zustand/vanilla";
import {
  ConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "@/modules/auth/domain/user";
import { AuthSrv } from "@/modules/auth/app/authSrv";

export interface AuthState {
  signedUpUser: SignedUpUser | null;
  connectedUser: ConnectedUserInfo | null;
  loading: boolean;
  error: string | null;

  signup: (userConnection: SignupUser) => Promise<boolean>;
  login: (userConnection: SigninUser) => Promise<boolean>;
  initialize: () => Promise<void>;
}

export const setAuthStore = (
  authSrv: AuthSrv,
  connectedUser?: ConnectedUserInfo,
) => {
  return createStore<AuthState>()((set, _) => ({
    signedUpUser: null,
    connectedUser: connectedUser ?? null,
    loading: false,
    error: null,

    signup: async (userConnection: SignupUser): Promise<boolean> => {
      set({
        loading: true,
        error: null,
      });

      try {
        const signedUpUser = await authSrv.signup(userConnection);
        set({
          signedUpUser: signedUpUser,
          loading: false,
        });
        return true;
      } catch (error) {
        set({
          loading: false,
          error: error instanceof Error ? error.message : "Unknown error",
        });
        return false;
      }
    },

    login: async (userConnection: SigninUser): Promise<boolean> => {
      set({
        loading: true,
        error: null,
      });

      try {
        const connectedUser = await authSrv.login(userConnection);
        set({
          connectedUser,
          loading: false,
        });
        return true;
      } catch (error) {
        set({
          connectedUser: null,
          loading: false,
          error: error instanceof Error ? error.message : "Unknown error",
        });
        return false;
      }
    },

    initialize: async () => {
        if (!authSrv.isTokenInMemory()){
          return
        }
        try{
          await authSrv.refresh()
          const user = await authSrv.whoami();
          console.log("-> user : ", user)
          set({
            connectedUser: user,
          });
        }catch(error){
          set({
            connectedUser: null,
          });
          throw Error("mlskdjf")
        }      
    },
  }));
};

export type AuthStoreApi = ReturnType<typeof setAuthStore>;
