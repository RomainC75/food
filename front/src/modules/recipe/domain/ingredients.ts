import { Ingredient, NewIngredient } from "./Ingredient"

export interface Ingredients {
    create(ingredient: NewIngredient): Promise<Ingredient>;
    getList(): Promise<Ingredient[]>;
    // update(ingredient: Ingredient): Promise<Ingredient>;
    // delete(workspace: string, id: string): Promise<void>;
}


