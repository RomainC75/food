import { createStore } from "zustand/vanilla";
import { Ingredient, NewIngredient } from "@/modules/recipe/domain/Ingredient";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";

export interface IngredientState {
  ingredients: Ingredient[];
  loading: boolean;
  error: string | null;

  fetchIngredients: () => Promise<void>;
}

export const setIngredientStore = (
  ingredientSrv: IngredientSrv
) => {
  return createStore<IngredientState>()((set) => ({
    ingredients: [],
    loading: false,
    error: null,

    fetchIngredients: async () => {
      set({
        loading: true,
        error: null,
      });

      try {
        console.log("---- fetch : ")
        const ingredients = await ingredientSrv.getList();

        set({
          ingredients,
          loading: false,
        });
      } catch (error) {
        set({
          loading: false,
          error: error instanceof Error
            ? error.message
            : "Unknown error",
        });
      }
    },

    saveIngredient: async (newIngredient: NewIngredient) => {
      set({
        loading: true,
        error: null,
      });

      try {
        const createdIngredient = await ingredientSrv.create(newIngredient);
        set((state)=>({
            ingredients: [...state.ingredients, createdIngredient],
            loading: false
        }))
      } catch (error) {
        set({
          loading: false,
          error: error instanceof Error
            ? error.message
            : "Unknown error",
        });
      }
    },
  }));
};

export type IngredientStoreApi = ReturnType<typeof setIngredientStore>;