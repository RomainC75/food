import { Ingredient, NewIngredient } from "../domain/Ingredient";
import { Ingredients } from "../domain/ingredients";

export class IngredientSrv {
    constructor(private ingredients: Ingredients){

    }

    create(ingredient: NewIngredient): Promise<Ingredient>{
        return this.ingredients.create(ingredient)
    }

    getList(): Promise<Ingredient[]>{
        return this.ingredients.getList()
    }
}