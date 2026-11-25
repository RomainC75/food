"use client";

import {
  createContext,
  useContext,
  ReactNode,
  useState,
} from "react";

import { useStore } from "zustand";
import { IngredientStoreApi, setIngredientStore } from "@/presentation/state-manager/stores/ingredient.store";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";

export const IngredientStoreContext = createContext<IngredientStoreApi | null>(null);

export function RecipeProvider({
  children,
  ingredientSrv,
}: {
  children: ReactNode;
  ingredientSrv: IngredientSrv;
}) {
  const [store] = useState(() => setIngredientStore(ingredientSrv));
  return (
    <IngredientStoreContext.Provider value={store}>
      {children}
    </IngredientStoreContext.Provider>
  );
}

export function useRecipeStore<T>(
  selector: (state: ReturnType<IngredientStoreApi["getState"]>) => T
): T {
  const store = useContext(IngredientStoreContext);

  if (!store) {
    throw new Error("RecipeStore provider is missing");
  }

  return useStore(store, selector);
}