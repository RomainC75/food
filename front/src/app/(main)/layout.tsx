"use client";
import ProtectedRoutesComp from "@/presentation/components/auth/ProtectedRoutes";
import MenuBar from "@/presentation/components/menuBar/MenuBar";

export default function MainLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <ProtectedRoutesComp>
      <div className="flex flex-col">
        <MenuBar />
        <main>{children}</main>
      </div>
    </ProtectedRoutesComp>
  );
}
