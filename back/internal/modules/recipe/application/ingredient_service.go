package recipe_app

import (
	"context"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	shared_domain "food/internal/shared/domain"
)

type IngredientSrv struct {
	uuidGen     shared_domain.Uuids
	ingredients recipe_dom.Ingredients
}

func NewIngredientSrv(uuidGen shared_domain.Uuids, ingredients recipe_dom.Ingredients) *IngredientSrv {
	return &IngredientSrv{
		uuidGen:     uuidGen,
		ingredients: ingredients,
	}
}

func (is *IngredientSrv) Create(ctx context.Context, ingredientReq dto_recipe_req.CreateIngredient) (recipe_dom.Ingredient, error) {

	ingredient := recipe_dom.NewIngredient(is.uuidGen.Generate(), ingredientReq.Name, ingredientReq.BaseUnit)

	err := is.ingredients.Save(ingredient)
	return ingredient, err
}
