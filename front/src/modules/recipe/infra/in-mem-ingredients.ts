import { Uuids } from "@/modules/shared/domain/uuid";
import { Ingredient, NewIngredient } from "../domain/Ingredient";
import { Ingredients } from "../domain/ingredients";
import { BaseUnit } from "../domain/baseUnit";

const ingredients: Ingredient[] = [
    new Ingredient("1", "salt", BaseUnit.Gram, "my-workspace"),
    new Ingredient("2", "pasta", BaseUnit.Gram, "my-workspace"),
]


export class InMemIngredients implements Ingredients{
    ingredients: Ingredient[];
    constructor(private uuidGen: Uuids){
        this.ingredients= ingredients
    }

    setIngredients(ingredients: Ingredient[]){
        this.ingredients = ingredients;
    }

    create(newIngredient: NewIngredient): Promise<Ingredient>{
        const ingredients = Ingredient.fromNewIngredient(this.uuidGen.now(), newIngredient)
        this.ingredients.push(ingredients)
        return Promise.resolve(ingredients)
    }

    getList(): Promise<Ingredient[]> {
        return Promise.resolve(this.ingredients)
    }
}