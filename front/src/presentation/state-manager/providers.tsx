import { RecipeProvider } from "./providers/recipeProvider";
import { AuthProvider } from "./providers/authProvider";
import { Dependencies } from "@/dependencies";

export default function Providers({
  children,
  deps,
}: {
  children: React.ReactNode;
  deps: Dependencies;
}) {
  return (
    <AuthProvider authSrv={deps.authSrv}>
      <RecipeProvider ingredientSrv={deps.ingredientSrv}>
        {children}
      </RecipeProvider>
    </AuthProvider>
  );
}
