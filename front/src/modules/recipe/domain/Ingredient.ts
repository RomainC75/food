import { BaseUnit } from "./BaseUnit";

export class NewIngredient {
  constructor(
    private name: string,
    private description: string,
    private baseUnit: BaseUnit, 
    private workspaceId: string
  ) {}
  

  setName(name: string): void {
    this.name = name;
  }

  setBaseUnit(baseUnit: BaseUnit): void {
    this.baseUnit = baseUnit;
  }

  isSameWorkspace(workspaceId: string): boolean {
    return this.workspaceId === workspaceId;
  }

  isQtyValid(qty: { unit: BaseUnit }): boolean {
    const family = Object.values(BaseUnit).find((units) => units.includes(this.baseUnit));
    return family ? family.includes(qty.unit) : true;
  }
}

export class Ingredient extends NewIngredient{
  constructor(
    private id: string, 
    name: string,
    description: string,
    baseUnit: BaseUnit, 
    workspaceId: string) {
      super(name, description, baseUnit, workspaceId)
  }

  get Id(): string{
    return this.id;
  }
}
