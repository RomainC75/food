import { createStore } from "zustand/vanilla";
import { ConnectedUserInfo, SignedUpUser, SigninUser, SignupUser } from "@/modules/auth/domain/user";
import { AuthSrv } from "@/modules/auth/app/authSrv";

export interface AuthState {
  signedUpUser: SignedUpUser | null;
  connectedUser: ConnectedUserInfo | null;
  loading: boolean;
  error: string | null;

  signup: (userConnection: SignupUser) => Promise<void>;
  login: (userConnection: SigninUser) => Promise<void>
}

export const setAuthStore = (
  authSrv: AuthSrv,
) => {
  return createStore<AuthState>()((set, _) => ({
    signedUpUser: null,
    connectedUser: null,
    loading: false,
    error: null,

    signup: async (userConnection: SignupUser) => {
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
      } catch (error) {
        set({
          loading: false,
          error: error instanceof Error
            ? error.message
            : "Unknown error",
        });
      }
    },

    login: async (userConnection: SigninUser) => {
      set({
        loading: true,
        error: null,
      });
      
      try {
        const connectedUser = await authSrv.login(userConnection);
        set(({
              connectedUser,
              loading: false
            })
        )
      } catch (error) {
        set({
          connectedUser: null,
          loading: false,
          error: error instanceof Error
            ? error.message
            : "Unknown error",
        });
      }
    },
  }));
};

export type AuthStoreApi = ReturnType<typeof setAuthStore>;