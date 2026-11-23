import { IngredientPlain } from "@/modules/recipe/init/ingredient.init";
import { RecipeProvider } from "./providers/recipeProvider";
import { AuthProvider } from "./providers/authProvider";

export default function TestProviders({
  children,
  options,
}: {
  children: React.ReactNode;
  options?: {
    initPlainIngredients?: IngredientPlain[];
    isUsersSlow?: boolean;
    isResponseError?: boolean;
    isUserUnverified?: boolean;
  };
}) {
  return (
    <AuthProvider
      isUsersSlow={!!options?.isUsersSlow}
      isResponseError={!!options?.isResponseError}
      isUserUnverified={!!options?.isUserUnverified}
    >
      <RecipeProvider
        initPlainIngredients={options?.initPlainIngredients ?? []}
      >
        {children}
      </RecipeProvider>
    </AuthProvider>
  );
}
