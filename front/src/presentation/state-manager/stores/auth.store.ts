import { createStore } from "zustand/vanilla";
import { ConnectedUserInfo, UserConnection } from "@/modules/auth/domain/user";
import { AuthSrv } from "@/modules/auth/app/authSrv";

export interface AuthState {
  connectedUser: ConnectedUserInfo | null
  loading: boolean;
  error: string | null;

  signup: (userConnection: UserConnection) => Promise<void>;
  login: (userConnection: UserConnection) => Promise<void>
}

export const setAuthStore = (
  authSrv: AuthSrv,
) => {
  return createStore<AuthState>()((set, _) => ({
    connectedUser: null,
    loading: false,
    error: null,

    signup: async (userConnection: UserConnection) => {
      set({
        loading: true,
        error: null,
      });

      try {
        await authSrv.signup(userConnection);
        set({
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

    login: async (userConnection: UserConnection) => {
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