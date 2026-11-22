import { BaseUnit } from "../domain/baseUnit";

export type IngredientPlain = {
    id: string,
    name: string,
    baseUnit: BaseUnit,
    workspaceId: string
}
export const initPlainIngredients: IngredientPlain[] = [
    {id:"1", name: "salt", baseUnit: BaseUnit.Gram, workspaceId: "my-workspace"},
    {id:"2", name: "pasta", baseUnit: BaseUnit.Gram, workspaceId: "my-workspace"},
];