"use client";
import ProtectedRoutesComp from "@/presentation/components/auth/ProtectedRoutes";
import HomePage from "@/presentation/pages/home/HomePage";

export default function Home() {
  return (
    <div className="flex flex-col flex-1 items-center justify-center bg-zinc-50 font-sans dark:bg-black">
        <HomePage />
    </div>
  );
}
