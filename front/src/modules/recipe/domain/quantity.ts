import { BaseUnit } from "./BaseUnit";

export class Quantity {
    static create(amount: number, unit: BaseUnit): Quantity{
        if (amount <= 0){
            throw new Error("Price cannot be negative");
        }
        return new Quantity(amount, unit);
    }

    private constructor(
        private amount: number, 
        private unit: BaseUnit
    ){}

    isQuantityCompatible(other: Quantity): boolean{
        if(this.unit == BaseUnit.Piece){
            if(other.unit != BaseUnit.Piece){
                return false
            }
        } else if (this.unit == BaseUnit.Gram || this.unit == BaseUnit.Kg){
            if (other.unit != BaseUnit.Gram && other.unit != BaseUnit.Kg){
                return false
            }
        } else if (this.unit == BaseUnit.Liter || this.unit == BaseUnit.Ml){
            if (other.unit != BaseUnit.Liter && other.unit != BaseUnit.Ml){
                return false
            }
        }
        return true
    }
}