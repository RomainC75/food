package dto_recipe_res

import "github.com/google/uuid"

type IngredientCreatedDTO struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	BaseUnit string    `json:"base_unit"`
}
