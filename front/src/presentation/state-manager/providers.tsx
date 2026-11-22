import { IngredientPlain } from "@/modules/recipe/init/ingredient.init";
import { RecipeProvider } from "./providers/recipeProvider";
import { AuthProvider } from "./providers/authProvider";


export function Providers({
  children,
  initPlainIngredients,
}: {
  children: React.ReactNode;
  initPlainIngredients: IngredientPlain[]
}) {
  return (
    <AuthProvider>
      <RecipeProvider initPlainIngredients={initPlainIngredients}>
          {children}
      </RecipeProvider>
    </AuthProvider>
  );
}
