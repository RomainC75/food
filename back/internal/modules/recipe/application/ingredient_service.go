package recipe_app

import (
	"context"
	"errors"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	platform_http "food/internal/platform/http"
	shared_domain "food/internal/shared/domain"

	"github.com/google/uuid"
)

var (
	ErrCouldNotGetWorkspaceId = "ingredient / could not get workdspaceId"
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

func (is *IngredientSrv) Create(ctx context.Context, ingredientReq dto_recipe_req.CreateIngredientDto) (*recipe_dom.Ingredient, error) {
	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(ctx)
	if !ok {
		return nil, errors.New(ErrCouldNotGetWorkspaceId)
	}

	ingredient := recipe_dom.NewIngredient(is.uuidGen.Generate(), ingredientReq.Name, ingredientReq.BaseUnit, workspaceId)

	err := is.ingredients.Save(ctx, ingredient)
	return ingredient, err
}

func (is *IngredientSrv) Search(ctx context.Context, f recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	return is.ingredients.Search(ctx, f)
}

func (is *IngredientSrv) Update(ctx context.Context, uir dto_recipe_req.UpdateIngredientDto) (*recipe_dom.Ingredient, error) {
	found, err := is.ingredients.GetById(ctx, uir.Id)
	if err != nil {
		return nil, errors.New(recipe_dom.ErrIngredientNotFound)
	}
	if uir.Name != "" {
		found.SetName(uir.Name)
	}
	if uir.BaseUnit != "" {
		found.SetBaseUnit(uir.BaseUnit)
	}
	err = is.ingredients.Update(ctx, found)
	return found, err
}

func (is *IngredientSrv) DeleteById(ctx context.Context, id uuid.UUID) error {
	_, err := is.ingredients.GetById(ctx, id)
	if err != nil {
		return err
	}

	return is.ingredients.DeleteById(ctx, id)

}
