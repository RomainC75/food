package recipe_app

import (
	"context"
	"errors"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	platform_http "food/internal/platform/http"
	shared_domain "food/internal/shared/domain"
	"log"

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
	workspaceId, _ := platform_http.GetWorkspaceIdFromContext(ctx)
	uuid := is.uuidGen.Generate()

	log.Println("------------- uuid + workspace ", uuid, workspaceId)
	ingredient := recipe_dom.NewIngredient(uuid, ingredientReq.Name, ingredientReq.BaseUnit, workspaceId)
	b, _ := ingredient.MarshalJSON()
	log.Println("----saved : ", string(b))
	err := is.ingredients.Save(ctx, ingredient)
	return ingredient, err
}

func (is *IngredientSrv) Search(ctx context.Context, f recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	return is.ingredients.Search(ctx, f)
}

func (is *IngredientSrv) GetById(ctx context.Context, ingredientId uuid.UUID) (*recipe_dom.Ingredient, error) {
	workspaceId, _ := platform_http.GetWorkspaceIdFromContext(ctx)
	return is.ingredients.GetById(ctx, workspaceId, ingredientId)

}

func (is *IngredientSrv) Update(ctx context.Context, uir dto_recipe_req.UpdateIngredientDto) (*recipe_dom.Ingredient, error) {
	workspaceId, _ := platform_http.GetWorkspaceIdFromContext(ctx)

	found, err := is.ingredients.GetById(ctx, workspaceId, uir.Id)
	if err != nil {
		return nil, errors.New(recipe_dom.ErrIngredientNotFound)
	}
	if uir.Name != "" {
		found.SetName(uir.Name)
	}
	if uir.BaseUnit != "" {
		found.SetBaseUnit(uir.BaseUnit)
	}
	err = is.ingredients.Update(ctx, workspaceId, found)
	return found, err
}

func (is *IngredientSrv) DeleteById(ctx context.Context, id uuid.UUID) error {
	workspaceId, _ := platform_http.GetWorkspaceIdFromContext(ctx)
	return is.ingredients.DeleteById(ctx, workspaceId, id)

}
