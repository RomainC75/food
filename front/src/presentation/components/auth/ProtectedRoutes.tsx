"use client";

import React, { useEffect } from "react";
import { usePathname, useRouter } from "next/navigation";
import { useShallow } from "zustand/shallow";
import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";
import FullPageSpinner from "@/presentation/components/shared/FullPageSpinner";
import { loginUrlFor } from "./redirect";

export { LOGIN_URL } from "./redirect";

type ProtectedRoutesCompProps = {
  children: React.ReactNode;
};

/** Renders children only for an authenticated user, redirects to login otherwise. */
const ProtectedRoutesComp = ({ children }: ProtectedRoutesCompProps) => {
  const router = useRouter();
  const pathname = usePathname();
  const { status, initialize } = useAuthStore(
    useShallow((state) => ({
      status: state.status,
      initialize: state.initialize,
    })),
  );

  useEffect(() => {
    initialize();
  }, [initialize]);

  useEffect(() => {
    if (status === "unauthenticated") {
      router.replace(loginUrlFor(pathname));
    }
  }, [status, pathname, router]);

  if (status !== "authenticated") return <FullPageSpinner />;
  return <>{children}</>;
};

export default ProtectedRoutesComp;
