import { createStore } from "zustand/vanilla";
import {
  ConnectedUserInfo,
  SignedUpUser,
  SigninUser,
  SignupUser,
} from "@/modules/auth/domain/user";
import { AuthSrv } from "@/modules/auth/app/authSrv";
import { WorkspaceSrv } from "@/modules/auth/app/workspaceSrv";
import {
  InvalidCredentialsError,
  UserUnverifiedError,
} from "@/modules/auth/domain/errors";

/**
 * unknown          -> session not checked yet
 * checking         -> restoring the session from stored tokens
 * authenticated    -> connectedUser is set
 * unauthenticated  -> no session: login required
 */
export type AuthStatus = "unknown" | "checking" | "authenticated" | "unauthenticated";
export type LoginError = "invalid_credentials" | "unverified" | "unknown";

export interface AuthState {
  status: AuthStatus;
  connectedUser: ConnectedUserInfo | null;
  /** Set when the session check failed for a reason other than an expired session (e.g. network). */
  sessionError: string | null;

  signedUpUser: SignedUpUser | null;
  signupPending: boolean;
  signupError: string | null;

  loginPending: boolean;
  loginError: LoginError | null;

  createWorkspacePending: boolean;
  createWorkspaceError: string | null;

  /** Idempotent: only the first call restores the session, later calls share its promise. */
  initialize: () => Promise<void>;
  signup: (user: SignupUser) => Promise<boolean>;
  login: (credentials: SigninUser) => Promise<boolean>;
  logout: () => void;
  createWorkspace: (name: string) => Promise<boolean>;
}

const errorMessage = (error: unknown) =>
  error instanceof Error ? error.message : "Unknown error";

const toLoginError = (error: unknown): LoginError => {
  if (error instanceof UserUnverifiedError) return "unverified";
  if (error instanceof InvalidCredentialsError) return "invalid_credentials";
  return "unknown";
};

export const setAuthStore = (
  authSrv: AuthSrv,
  workspaceSrv: WorkspaceSrv,
  connectedUser?: ConnectedUserInfo,
) => {
  let initializing: Promise<void> | null = null;

  const store = createStore<AuthState>()((set, get) => ({
    status: connectedUser ? "authenticated" : "unknown",
    connectedUser: connectedUser ?? null,
    sessionError: null,

    signedUpUser: null,
    signupPending: false,
    signupError: null,

    loginPending: false,
    loginError: null,

    createWorkspacePending: false,
    createWorkspaceError: null,

    initialize: () => {
      if (initializing) return initializing;
      if (get().status !== "unknown") return Promise.resolve();

      set({ status: "checking", sessionError: null });
      initializing = (async () => {
        try {
          const user = await authSrv.restoreSession();
          // A login may have completed meanwhile: don't overwrite it
          if (get().status !== "checking") return;
          set(
            user
              ? { status: "authenticated", connectedUser: user }
              : { status: "unauthenticated", connectedUser: null },
          );
        } catch (error) {
          if (get().status !== "checking") return;
          set({
            status: "unauthenticated",
            connectedUser: null,
            sessionError: errorMessage(error),
          });
        }
      })();
      return initializing;
    },

    signup: async (user) => {
      set({ signupPending: true, signupError: null });
      try {
        const signedUpUser = await authSrv.signup(user);
        set({ signedUpUser, signupPending: false });
        return true;
      } catch (error) {
        set({ signupPending: false, signupError: errorMessage(error) });
        return false;
      }
    },

    login: async (credentials) => {
      set({ loginPending: true, loginError: null });
      try {
        const user = await authSrv.login(credentials);
        set({
          status: "authenticated",
          connectedUser: user,
          sessionError: null,
          loginPending: false,
        });
        return true;
      } catch (error) {
        set({ loginPending: false, loginError: toLoginError(error) });
        return false;
      }
    },

    logout: () => {
      authSrv.logout();
      set({ status: "unauthenticated", connectedUser: null });
    },

    createWorkspace: async (name) => {
      set({ createWorkspacePending: true, createWorkspaceError: null });
      try {
        const workspace = await workspaceSrv.create(name);
        set((state) => ({
          createWorkspacePending: false,
          connectedUser: state.connectedUser && {
            ...state.connectedUser,
            workspaces: [...state.connectedUser.workspaces, workspace],
          },
        }));
        return true;
      } catch (error) {
        set({
          createWorkspacePending: false,
          createWorkspaceError: errorMessage(error),
        });
        return false;
      }
    },
  }));

  // Session lost outside of a user action (refresh rejected, logout in another tab)
  authSrv.onSessionExpired(() =>
    store.setState({ status: "unauthenticated", connectedUser: null }),
  );

  return store;
};

export type AuthStoreApi = ReturnType<typeof setAuthStore>;
