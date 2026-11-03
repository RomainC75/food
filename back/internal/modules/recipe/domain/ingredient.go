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
