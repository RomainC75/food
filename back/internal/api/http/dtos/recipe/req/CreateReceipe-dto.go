package dto_recipe_req

import (
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/google/uuid"
)

type IngredientBasicInfoDto struct {
	Id   uuid.UUID           `json:"id" validate:"required"`
	Qty  float32             `json:"qty" validate:"required,gt=0"`
	Unit recipe_dom.BaseUnit `json:"base_unit" validate:"required,oneof=piece l ml kg g"`
}
type CreateRecipeDto struct {
	Name        string                   `json:"name" validate:"required"`
	Ingredients []IngredientBasicInfoDto `json:"ingredients" validate:"required"`
}

type UpdateRecipeDto struct {
	Id uuid.UUID
	CreateRecipeDto
}
