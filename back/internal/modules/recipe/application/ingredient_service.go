package recipe_app

import (
	"context"
	"errors"
	recipe_dom "food/internal/modules/recipe/domain"
	shared_domain "food/internal/shared/domain"

	"github.com/google/uuid"
)

var (
	ErrCouldNotGetWorkspaceId  = "ingredient / could not get workdspaceId"
	ErrIngredientAlreadyExists = "ingredient already exists"
)

type CreateIngredientCmd struct {
	WorkspaceId uuid.UUID
	Name        string
	BaseUnit    recipe_dom.BaseUnit
}

type GetIngredientByIdCmd struct {
	WorkspaceId  uuid.UUID
	IngredientId uuid.UUID
}
type UpdateIngredientCmd struct {
	WorkspaceId uuid.UUID
	Id          uuid.UUID
	CreateIngredientCmd
}

type DeleteIngredientCmd struct {
	WorkspaceId uuid.UUID
	Id          uuid.UUID
}

// ======================

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

func (is *IngredientSrv) Create(ctx context.Context, cmd CreateIngredientCmd) (*recipe_dom.Ingredient, error) {
	exists, err := is.ingredients.ExistsByName(ctx, cmd.WorkspaceId, cmd.Name)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New(ErrIngredientAlreadyExists)
	}

	uuid := is.uuidGen.Generate()
	ingredient := recipe_dom.NewIngredient(uuid, cmd.Name, cmd.BaseUnit, cmd.WorkspaceId)
	err = is.ingredients.Save(ctx, ingredient)
	return ingredient, err
}

func (is *IngredientSrv) Search(ctx context.Context, workspaceId uuid.UUID, f recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	return is.ingredients.Search(ctx, workspaceId, f)
}

func (is *IngredientSrv) GetById(ctx context.Context, cmd GetIngredientByIdCmd) (*recipe_dom.Ingredient, error) {
	return is.ingredients.GetById(ctx, cmd.WorkspaceId, cmd.IngredientId)

}

func (is *IngredientSrv) Update(ctx context.Context, cmd UpdateIngredientCmd) (*recipe_dom.Ingredient, error) {
	found, err := is.ingredients.GetById(ctx, cmd.WorkspaceId, cmd.Id)
	if err != nil {
		return nil, err
	}
	if cmd.Name != "" {
		found.SetName(cmd.Name)
	}
	if cmd.BaseUnit != "" {
		found.SetBaseUnit(cmd.BaseUnit)
	}
	err = is.ingredients.Update(ctx, cmd.WorkspaceId, found)
	return found, err
}

func (is *IngredientSrv) DeleteById(ctx context.Context, cmd DeleteIngredientCmd) error {
	return is.ingredients.DeleteById(ctx, cmd.WorkspaceId, cmd.Id)

}
