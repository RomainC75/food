import { IngredientPlain } from "@/modules/recipe/init/ingredient.init";
import { RecipeProvider } from "./providers/recipeProvider";
import { AuthProvider } from "./providers/authProvider";
import { AuthSrv } from "@/modules/auth/app/authSrv";

export default function Providers({
  children,
  authSrv,
  options
}: {
  children: React.ReactNode;
  authSrv: AuthSrv
  options?: {
    initPlainIngredients?: IngredientPlain[];
  };
}) {
  return (
    <AuthProvider
      authSrv={authSrv}
    >
      <RecipeProvider
        initPlainIngredients={options?.initPlainIngredients ?? []}
      >
        {children}
      </RecipeProvider>
    </AuthProvider>
  );
}
