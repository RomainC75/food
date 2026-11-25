import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";
import { useRouter } from "next/navigation";
import React, { useEffect } from "react";
import { useShallow } from "zustand/shallow";

export const LOGIN_URL = "/auth/login"

type ProtectedRoutesCompProps = {
  children: React.ReactNode;
};

const ProtectedRoutesComp = ({ children }: ProtectedRoutesCompProps) => {
  const router = useRouter();
  const { isUserConnected } = useAuthStore(
    useShallow((state) => ({
      isUserConnected: !!state.connectedUser,
    })),
  );

  // avoid infinit loop with it !
  useEffect(()=>{
    if (!isUserConnected) {
      router.replace(LOGIN_URL);
    }
  }, [isUserConnected, router])

  if (!isUserConnected) return null;
  return <>{children}</>;
};

export default ProtectedRoutesComp;
