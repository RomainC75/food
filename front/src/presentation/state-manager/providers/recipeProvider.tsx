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
import { InMemIngredients } from "@/modules/recipe/infra/in-mem-ingredients";
import { InMemUuidGen } from "@/modules/shared/infra/in-mem-uuidGen";
import { Ingredient } from "@/modules/recipe/domain/Ingredient";
import { IngredientPlain } from "@/modules/recipe/init/ingredient.init";

export const IngredientStoreContext = createContext<IngredientStoreApi | null>(null);

export function RecipeProvider({
  children, 
  initPlainIngredients = [],
}: {
  children: ReactNode;
  initPlainIngredients?: IngredientPlain[];
}) {
  const [store] = useState(() => {
    const uuidGen = new InMemUuidGen();
    const repository = new InMemIngredients(uuidGen, initPlainIngredients.map(ip => 
      new Ingredient(ip.id, ip.name, ip.baseUnit, ip.workspaceId)));
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