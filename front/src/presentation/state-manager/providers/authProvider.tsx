"use client";

import { createContext, useContext, ReactNode, useState } from "react";
import { useStore } from "zustand";
import { AuthSrv } from "@/modules/auth/app/authSrv";
import { WorkspaceSrv } from "@/modules/auth/app/workspaceSrv";
import {
  AuthState,
  AuthStoreApi,
  setAuthStore,
} from "@/presentation/state-manager/stores/auth.store";
import { ConnectedUserInfo } from "@/modules/auth/domain/user";

export const AuthStoreContext = createContext<AuthStoreApi | null>(null);

export function AuthProvider({
  children,
  authSrv,
  workspaceSrv,
  connectedUser,
}: {
  children: ReactNode;
  authSrv: AuthSrv;
  workspaceSrv: WorkspaceSrv;
  connectedUser?: ConnectedUserInfo;
}) {
  const [store] = useState(() => setAuthStore(authSrv, workspaceSrv, connectedUser));
  return (
    <AuthStoreContext.Provider value={store}>
      {children}
    </AuthStoreContext.Provider>
  );
}

export function useAuthStore<T>(selector: (state: AuthState) => T): T {
  const store = useContext(AuthStoreContext);
  if (!store) {
    throw new Error("useAuthStore must be used inside <AuthProvider>");
  }
  return useStore(store, selector);
}
