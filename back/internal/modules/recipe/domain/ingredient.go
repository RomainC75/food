package recipe_dom

import "github.com/google/uuid"

type Ingredient struct {
	id       uuid.UUID
	name     string
	baseUnit BaseUnit
}

func NewIngredient(id uuid.UUID, name string, baseUnit BaseUnit) *Ingredient {
	return &Ingredient{
		id:       id,
		name:     name,
		baseUnit: baseUnit,
	}
}

func (ing *Ingredient) Id() uuid.UUID {
	return ing.id
}

func (ing *Ingredient) Name() string {
	return ing.name
}

func (ing *Ingredient) BaseUnit() BaseUnit {
	return ing.baseUnit
}

func (ing *Ingredient) IsQtyValid(q Quantity) bool {
	if ing.baseUnit == Piece {
		if q.unit != Piece {
			return false
		}
	}
	if ing.baseUnit == Liter || ing.baseUnit == MilliLiter {
		if q.unit != Liter && q.unit != MilliLiter {
			return false
		}
	}
	if ing.baseUnit == Kilo || ing.baseUnit == Gram {
		if q.unit != Kilo && q.unit != Gram {
			return false
		}
	}
	return true
}
