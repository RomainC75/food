import { Uuids } from "@/modules/shared/domain/uuid";
import { Ingredient, NewIngredient } from "../domain/Ingredient";
import { Ingredients } from "../domain/ingredients";
import { BaseUnit } from "../domain/baseUnit";



export class InMemIngredients implements Ingredients{
    savedIngredients: Ingredient[];
    constructor(private uuidGen: Uuids, initPlainIngredients: Ingredient[]){
        this.savedIngredients= initPlainIngredients;
        this.savedIngredients = initPlainIngredients;
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