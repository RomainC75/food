"use client";

import {
  createContext,
  useContext,
  ReactNode,
  useState,
} from "react";

import { useStore } from "zustand";
import { AuthSrv } from "@/modules/auth/app/authSrv";
import { AuthStoreApi, setAuthStore } from "@/presentation/state-manager/stores/auth.store";

export const AuthStoreContext = createContext<AuthStoreApi | null>(null);

export function AuthProvider({
  children, 
  authSrv
}: {
  children: ReactNode;
  authSrv: AuthSrv;
}) {
  const [store] = useState(() => {
    return setAuthStore(authSrv);
  });
  return (
    <AuthStoreContext.Provider value={store}>
      {children}
    </AuthStoreContext.Provider>
  );
}

export function useAuthStore<T>(
  selector: (state: ReturnType<AuthStoreApi["getState"]>) => T
): T {
  const store = useContext(AuthStoreContext);

  if (!store) {
    throw new Error("test AuthStore provider is missing");
  }

  return useStore(store, selector);
}