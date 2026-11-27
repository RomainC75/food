"use client";

import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useShallow } from "zustand/shallow";
import { useAuthStore } from "@/presentation/state-manager/providers/authProvider";
import FullPageSpinner from "@/presentation/components/shared/FullPageSpinner";
import { nextPathFromLocation } from "./redirect";

type GuestRoutesCompProps = {
  children: React.ReactNode;
};

/** For login/signup: an already authenticated user is sent away. */
const GuestRoutesComp = ({ children }: GuestRoutesCompProps) => {
  const router = useRouter();
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
    if (status === "authenticated") {
      router.replace(nextPathFromLocation());
    }
  }, [status, router]);

  if (status === "unauthenticated") return <>{children}</>;
  return <FullPageSpinner />;
};

export default GuestRoutesComp;
