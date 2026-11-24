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
import { Users } from "@/modules/auth/domain/users";
import { HttpUsers } from "@/modules/auth/infra/http-users";

export const AuthStoreContext = createContext<AuthStoreApi | null>(null);

export function AuthProvider({
  children, 
  isForTest,
  isUsersSlow,
  isResponseError,
  isUserUnverified,
}: {
  children: ReactNode;
  isForTest?: boolean,
  isUsersSlow?: boolean,
  isResponseError?: boolean,
  isUserUnverified?: boolean,
}) {
  const [store] = useState(() => {
    let users: Users;
    if(isForTest){
      const uuids = new InMemUuidGen()
      users = new InMemUser(uuids, !!isUsersSlow)
      .setIsResponseError(!!isResponseError)
      .setIsUserUnverified(!!isUserUnverified)
    }else {
      users = new HttpUsers()
    }

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
  console.log("-----> store : ", store)

  if (!store) {
    throw new Error("test AuthStore provider is missing");
  }

  return useStore(store, selector);
}