import { Ingredient, NewIngredient } from "./Ingredient"

export interface Ingredients {
    Create(ingredient: NewIngredient): Promise<Ingredient>;
    GetList(workspace: string): Promise<Ingredient[]>;
    Update(ingredient: Ingredient): Promise<Ingredient>;
    Delete(workspace: string, id: string): Promise<void>;
}


