import { Uuids } from "@/modules/shared/domain/uuid";
import { Ingredient, NewIngredient } from "../domain/Ingredient";
import { Ingredients } from "../domain/ingredients";


export class InMemIngredient implements Ingredients{
    ingredients: Ingredient[];
    constructor(private uuidGen: Uuids){
        this.ingredients= []
    }

    setIngredients(ingredients: Ingredient[]){
        this.ingredients = ingredients;
    }

    create(newIngredient: NewIngredient): Promise<Ingredient>{
        const ingredients = Ingredient.fromNewIngredient(this.uuidGen.now(), newIngredient)
        this.ingredients.push(ingredients)
        return Promise.resolve(ingredients)
    }
}