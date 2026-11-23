import { RecipeProvider } from "./providers/recipeProvider";
import { AuthProvider } from "./providers/authProvider";


export default function Providers({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <AuthProvider>
      <RecipeProvider>
          {children}
      </RecipeProvider>
    </AuthProvider>
  );
}
