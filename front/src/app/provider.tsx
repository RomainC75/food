"use client";

import {
  createContext,
  useContext,
  useRef,
  ReactNode,
  useState,
} from "react";

import { useStore } from "zustand";
import { IngredientStoreApi, setIngredientStore } from "@/presentation/stores/ingredient.store";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";
import { InMemIngredients } from "@/modules/recipe/infra/in-mem-ingredients";
import { UuidGen } from "@/modules/shared/infra/uuidGen";

const IngredientStoreContext =
  createContext<IngredientStoreApi | null>(null);

export function Providers({
  children,
}: {
  children: ReactNode;
}) {
  // const storeRef = useRef<IngredientStoreApi | null>(null);
  // if (!storeRef.current) {
  //   // ! don't forger to setup the server url
  //   const uuidGen = new UuidGen()
  //   const repository = new InMemIngredients(uuidGen);
  //   const ingredientSrv = new IngredientSrv(repository);

  //   storeRef.current = setIngredientStore(ingredientSrv);
  // }
  const [store] = useState(() => {
  const uuidGen = new UuidGen();
  const repository = new InMemIngredients(uuidGen);
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