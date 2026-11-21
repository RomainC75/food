import { Ingredient } from "./Ingredient";

export class NewRecipe {
  constructor(private workspace: string, private name: string, private ingredients: Ingredient[]){
  }

  public addIngredient(ingredient: Ingredient){
    const foundIndex = this.ingredients.findIndex(i => i.Id === ingredient.Id);
    if (foundIndex >= 0 ){
        this.ingredients[foundIndex] = ingredient
    }else {
        this.ingredients.push(ingredient)
    }
  }
};

export class Recipe extends NewRecipe{
    constructor(private id: string, workspace: string, name: string, ingredients: Ingredient[]){
        super(workspace, name, ingredients);
    }
}