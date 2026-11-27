import { RecipeProvider } from "./providers/recipeProvider";
import { AuthProvider } from "./providers/authProvider";
import { Dependencies } from "@/dependencies";
import { ConnectedUserInfo } from "@/modules/auth/domain/user";

export default function Providers({
  children,
  deps,
  connectedUser,
}: {
  children: React.ReactNode;
  deps: Dependencies;
  connectedUser?: ConnectedUserInfo;
}) {
  return (
    <AuthProvider
      authSrv={deps.authSrv}
      workspaceSrv={deps.workspaceSrv}
      connectedUser={connectedUser}
    >
      <RecipeProvider ingredientSrv={deps.ingredientSrv}>
        {children}
      </RecipeProvider>
    </AuthProvider>
  );
}
