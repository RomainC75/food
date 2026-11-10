package recipe_dom

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Ingredient struct {
	id          uuid.UUID
	name        string
	baseUnit    BaseUnit
	workspaceId uuid.UUID
}

func NewIngredient(id uuid.UUID, name string, baseUnit BaseUnit, workspaceId uuid.UUID) *Ingredient {
	return &Ingredient{
		id:          id,
		name:        name,
		baseUnit:    baseUnit,
		workspaceId: workspaceId,
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

func (ing *Ingredient) SetName(new string) {
	ing.name = new
}

func (ing *Ingredient) SetBaseUnit(new BaseUnit) {
	ing.baseUnit = new
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

func (i *Ingredient) Clone() *Ingredient {
	return &Ingredient{
		id:       i.id,
		name:     i.name,
		baseUnit: i.baseUnit,
	}
}

func (i *Ingredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		BaseUnit BaseUnit  `json:"base_unit"`
	}{
		ID:       i.id,
		Name:     i.name,
		BaseUnit: i.baseUnit,
	})
}
