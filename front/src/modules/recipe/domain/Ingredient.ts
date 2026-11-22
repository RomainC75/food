import { BaseUnit } from "./baseUnit";

export class NewIngredient {
  constructor(
    private name: string,
    private baseUnit: BaseUnit, 
    private workspaceId: string
  ) {}
  
  get Name(): string{
    return this.name;
  }


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

  export(){
    return {
      name: this.name,
      baseUnit: this.baseUnit,
      workspaceId: this.workspaceId,
    }
  }
}

export class Ingredient extends NewIngredient{
  constructor(
    private id: string, 
    name: string,
    baseUnit: BaseUnit, 
    workspaceId: string) {
      super(name, baseUnit, workspaceId)
  }

  get Id(): string{
    return this.id;
  }

  static fromNewIngredient(id: string, newIngredient: NewIngredient): Ingredient{
    const {
      name, baseUnit, workspaceId
    } = newIngredient.export()
    return new Ingredient(id, name, baseUnit, workspaceId);
  }
}
