"use client";
import GuestRoutesComp from "@/presentation/components/auth/GuestRoutes";

export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return <GuestRoutesComp>{children}</GuestRoutesComp>;
}
