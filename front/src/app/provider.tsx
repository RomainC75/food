"use client";

import {
  createContext,
  useContext,
  ReactNode,
  useState,
} from "react";

import { useStore } from "zustand";
import { IngredientStoreApi, setIngredientStore } from "@/presentation/stores/ingredient.store";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";
import { InMemIngredients } from "@/modules/recipe/infra/in-mem-ingredients";
import { InMemUuidGen } from "@/modules/shared/infra/in-mem-uuidGen";
import { Ingredient } from "@/modules/recipe/domain/Ingredient";

export const IngredientStoreContext = createContext<IngredientStoreApi | null>(null);

export function Providers({
  children, 
  initIngredients = [],
}: {
  children: ReactNode;
  initIngredients: Ingredient[];
}) {
  const [store] = useState(() => {
    const uuidGen = new InMemUuidGen();
    const repository = new InMemIngredients(uuidGen, initIngredients);
    const ingredientSrv = new IngredientSrv(repository);

    return setIngredientStore(ingredientSrv);
  });
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