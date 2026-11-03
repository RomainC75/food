package dto_recipe_req

import (
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/google/uuid"
)

type IngredientBasicInfo struct {
	Id   uuid.UUID
	Qty  float32
	Unit recipe_dom.BaseUnit
}

type CreateRecipeDto struct {
	Name        string
	Ingredients []IngredientBasicInfo
}

type UpdateRecipeDto struct {
	Id uuid.UUID
	CreateRecipeDto
}
