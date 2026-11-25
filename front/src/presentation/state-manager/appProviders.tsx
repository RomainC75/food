"use client";

import { useState } from "react";
import Providers from "./providers";
import { createAppDependencies } from "@/dependencies";

// Builds the real dependencies once, on the client
// (class instances can't be passed from a Server Component like layout.tsx)
export default function AppProviders({ children }: { children: React.ReactNode }) {
  const [deps] = useState(createAppDependencies);

  return <Providers deps={deps}>{children}</Providers>;
}
