import { Ingredient, NewIngredient } from "../domain/Ingredient";
import { Ingredients } from "../domain/ingredients";

export class IngredientSrv {
    constructor(private ingredients: Ingredients){

    }

    createIngredient(ingredient: NewIngredient): Promise<Ingredient>{
        return this.ingredients.create(ingredient)
    }
}