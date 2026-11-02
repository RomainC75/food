package dto_recipe_req

import recipe_dom "food/internal/modules/recipe/domain"

type CreateIngredient struct {
	Name     string              `json:"name" validate:"required"`
	BaseUnit recipe_dom.BaseUnit `json:"base_unit" validate:"required,oneof=piece l mk kg g"`
}
