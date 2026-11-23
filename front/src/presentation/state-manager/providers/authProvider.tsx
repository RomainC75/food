"use client";

import {
  createContext,
  useContext,
  ReactNode,
  useState,
} from "react";

import { useStore } from "zustand";
import { InMemUuidGen } from "@/modules/shared/infra/in-mem-uuidGen";
import { AuthSrv } from "@/modules/auth/app/authSrv";
import { InMemUser } from "@/modules/auth/infra/in-mem-users";
import { AuthStoreApi, setAuthStore } from "@/presentation/state-manager/stores/auth.store";

export const AuthStoreContext = createContext<AuthStoreApi | null>(null);

export function AuthProvider({
  children, 
  isUsersSlow,
  isResponseError
}: {
  children: ReactNode;
  isUsersSlow?: boolean,
  isResponseError?: boolean,
}) {
  const [store] = useState(() => {
    const uuids = new InMemUuidGen()
    const users = new InMemUser(uuids, !!isUsersSlow).setIsResponseError(!!isResponseError)
    const authSrv = new AuthSrv(users)
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
    throw new Error("AuthStore provider is missing");
  }

  return useStore(store, selector);
}