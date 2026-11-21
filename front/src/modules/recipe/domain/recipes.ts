import { NewRecipe, Recipe } from "./recipe"

export interface Recipes {
    Create(recipe: NewRecipe): Promise<Recipe>
    GetList(workspace: string): Promise<Recipes[]>
    Update(recipe: Recipe): Promise<Recipe>
    Delete(id: string): Promise<void>
}



