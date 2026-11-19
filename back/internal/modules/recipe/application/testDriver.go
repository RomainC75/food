package recipe_app

import (
	"context"
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	shared_infra "food/internal/shared/infra"
	"food/utils/displayer"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestDriver struct {
	t                *testing.T
	ctx              context.Context
	inMemUuidGen     *shared_infra.InMemUuidGen
	inMemIngredients *recipe_infra.InMemoryIngredients
	inMemRecipes     *recipe_infra.InMemRecipes
	ingredientSrv    *IngredientSrv
	recipeSrv        *RecipeSrv
}

func NewTestDriver(t *testing.T) *TestDriver {
	inMemUuidGen := shared_infra.NewInMemUuidGen()
	inMemIngredients := recipe_infra.NewInMemoryIngredients()
	inMemRecipes := recipe_infra.NewInMemRecipes()
	ingredientSrv := NewIngredientSrv(inMemUuidGen, inMemIngredients)
	recipeSrv := NewRecipeSrv(inMemUuidGen, inMemIngredients, inMemRecipes)

	return &TestDriver{
		t:                t,
		ctx:              context.Background(),
		inMemUuidGen:     inMemUuidGen,
		inMemIngredients: inMemIngredients,
		inMemRecipes:     inMemRecipes,
		ingredientSrv:    ingredientSrv,
		recipeSrv:        recipeSrv,
	}
}

func (td *TestDriver) SetNextUuid(id uuid.UUID) *TestDriver {
	td.inMemUuidGen.SetNext(id)
	return td
}

func (td *TestDriver) SetNextUuids(ids ...uuid.UUID) *TestDriver {
	td.inMemUuidGen.SetMultipleUuids(ids)
	return td
}

func (td *TestDriver) CreateIngredient(workspaceid uuid.UUID, cmd CreateIngredientCmd) (*recipe_dom.Ingredient, error) {
	return td.ingredientSrv.Create(td.ctx, cmd)
}

func (td *TestDriver) DeleteIngredient(workspaceId uuid.UUID, id uuid.UUID) error {
	deleteCmd := DeleteIngredientCmd{
		Id:          id,
		WorkspaceId: workspaceId,
	}
	return td.ingredientSrv.DeleteById(td.ctx, deleteCmd)
}

func (td *TestDriver) GivenIngredients(ings ...*recipe_dom.Ingredient) *TestDriver {
	td.t.Helper()
	for _, ing := range ings {
		require.NoError(td.t, td.inMemIngredients.Save(td.ctx, ing))
	}
	return td
}

func (td *TestDriver) GetDeletedLastIngredient() *recipe_dom.Ingredient {
	return td.inMemIngredients.GetLastDeleted()
}

func (td *TestDriver) GetIngredientById(workspaceId uuid.UUID, cmd GetIngredientByIdCmd) (*recipe_dom.Ingredient, error) {
	return td.ingredientSrv.GetById(td.ctx, cmd)
}

func (td *TestDriver) Search(workspaceId uuid.UUID, workdspaceId uuid.UUID, i recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	return td.ingredientSrv.Search(td.ctx, workdspaceId, i)
}

func (td *TestDriver) UpdateIngredient(workspaceId uuid.UUID, cmd UpdateIngredientCmd) (*recipe_dom.Ingredient, error) {
	return td.ingredientSrv.Update(td.ctx, cmd)
}

// -- RECEIPE --
// ?? expectedId ??
func (td *TestDriver) GivenRecipe(expectedId uuid.UUID, cmd CreateRecipeCmd) *recipe_dom.Recipe {
	td.t.Helper()
	td.SetNextUuid(expectedId)
	r, err := td.CreateRecipe(cmd)
	require.NoError(td.t, err)
	return r
}

func (td *TestDriver) CreateRecipe(cmd CreateRecipeCmd) (*recipe_dom.Recipe, error) {
	displayer.PrettyDisplay(" CREATE RECIPE CMD: ", cmd)
	return td.recipeSrv.Create(td.ctx, cmd)
}

func (td *TestDriver) UpdateRecipe(cmd UpdateRecipeCmd) (*recipe_dom.Recipe, error) {
	return td.recipeSrv.Update(td.ctx, cmd)
}

func (td *TestDriver) DeleteRecipe(workspaceId uuid.UUID, id uuid.UUID) error {
	return td.recipeSrv.Delete(td.ctx, workspaceId, id)
}

func (td *TestDriver) GivenRepoErrors(ets recipe_infra.RecipeRepoErrorsToSet) {
	td.inMemRecipes.SetExpectedErrors(ets)
}

func (td *TestDriver) GivenIngredientRepoErrors(expectedGetError bool, expectedUpdateError bool) {
	td.inMemIngredients.SetExpectedErrors(expectedGetError, expectedUpdateError)
}

// Verifiyer

func (td *TestDriver) SavedIngredients() map[uuid.UUID]*recipe_dom.Ingredient {
	return td.inMemIngredients.GetSaved()
}

func (td *TestDriver) SavedRecipes() []*recipe_dom.Recipe {
	return td.inMemRecipes.GetSavedRecipes()
}

func assertErrContainsAll(t *testing.T, err error, errorStrs ...string) {
	for _, errorsStr := range errorStrs {
		assert.Contains(t, err.Error(), errorsStr)
	}
}
