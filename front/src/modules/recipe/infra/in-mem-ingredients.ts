import { Uuids } from "@/modules/shared/domain/uuid";
import { Ingredient, NewIngredient } from "../domain/Ingredient";
import { Ingredients } from "../domain/ingredients";
import { BaseUnit } from "../domain/baseUnit";

const initIngredients: Ingredient[] = [
    new Ingredient("1", "salt", BaseUnit.Gram, "my-workspace"),
    new Ingredient("2", "pasta", BaseUnit.Gram, "my-workspace"),
];

export class InMemIngredients implements Ingredients{
    savedIngredients: Ingredient[];
    constructor(private uuidGen: Uuids){
        this.savedIngredients= initIngredients;
    }

    setIngredients(ingredients: Ingredient[]){
        this.savedIngredients = ingredients;
    }

    create(newIngredient: NewIngredient): Promise<Ingredient>{
        const ingredient = Ingredient.fromNewIngredient(this.uuidGen.now(), newIngredient)
        this.savedIngredients.push(ingredient)
        return Promise.resolve(ingredient)
    }

    getList(): Promise<Ingredient[]> {
        return Promise.resolve([...this.savedIngredients])
    }
}