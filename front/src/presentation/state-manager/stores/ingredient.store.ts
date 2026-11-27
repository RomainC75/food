import { createStore } from "zustand/vanilla";
import { Ingredient, NewIngredient } from "@/modules/recipe/domain/Ingredient";
import { IngredientSrv } from "@/modules/recipe/app/ingredientSrv";
import { CreateIngredientForm } from "../../components/recipe/createIngredient/create-ingredient";
import { Recipe } from "@/modules/recipe/domain/recipe";

export interface RecipeState {
  ingredients: Ingredient[];
  recipes: Recipe[];
  loading: boolean;
  error: string | null;
  workspace: string | null;

  fetchIngredients: () => Promise<void>;
  saveIngredient: (newIngredient: CreateIngredientForm) => Promise<void>
}

export const setIngredientStore = (
  ingredientSrv: IngredientSrv,
) => {
  return createStore<RecipeState>()((set, get) => ({
    ingredients: [],
    recipes: [],
    loading: false,
    error: null,
    workspace: null,

    fetchIngredients: async () => {
      set({
        loading: true,
        error: null,
      });

      try {
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

    saveIngredient: async ({name, baseUnit}: CreateIngredientForm) => {
      const { workspace } = get();
      if(!workspace){
        return
      }
      const newIngredient = new NewIngredient(name, baseUnit, workspace)
      set({
        loading: true,
        error: null,
      });
      
      try {
        const createdIngredient = await ingredientSrv.create(newIngredient);
        set((state)=>({
              ingredients: [...state.ingredients, createdIngredient],
              loading: false
            })
        )
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