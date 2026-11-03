package recipe_app

import (
	"context"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	shared_infra "food/internal/shared/infra"

	"github.com/google/uuid"
)

type TestDriver struct {
	ctx              context.Context
	inMemIngredients *recipe_infra.InMemoryIngredients
	inMemUuidGen     *shared_infra.InMemUuidGen
	ingredientSrv    *IngredientSrv
}

func NewTestDriver() *TestDriver {
	inMemUuidGen := shared_infra.NewInMemUuidGen()
	inMemIngredients := recipe_infra.NewInMemoryIngredients()
	ingredientSrv := NewIngredientSrv(inMemUuidGen, inMemIngredients)
	return &TestDriver{
		ctx:              context.Background(),
		inMemIngredients: inMemIngredients,
		inMemUuidGen:     inMemUuidGen,
		ingredientSrv:    ingredientSrv,
	}
}

func (td *TestDriver) SetNextUuid(id uuid.UUID) {
	td.inMemUuidGen.AddExpected(id)
}

func (td *TestDriver) CreateIngredient(ingredientDto dto_recipe_req.CreateIngredient) (recipe_dom.Ingredient, error) {
	return td.ingredientSrv.Create(td.ctx, ingredientDto)
}

func (td *TestDriver) GetSavedIngredients() []recipe_dom.Ingredient {
	return td.inMemIngredients.GetSaved()
}

func (td *TestDriver) DeleteIngredient(id uuid.UUID) error {
	return td.ingredientSrv.DeleteById(td.ctx, id)
}

func (td *TestDriver) SetIngredientInMemory(ingredient recipe_dom.Ingredient) {
	td.inMemIngredients.Save(td.ctx, ingredient)
}

func (td *TestDriver) GetDeletedLastIngredient() recipe_dom.Ingredient {
	return td.inMemIngredients.GetLastDeleted()
}

func (td *TestDriver) Search(i recipe_dom.IngredientFilter) (recipe_dom.Ingredient, error) {
	return td.ingredientSrv.Search(td.ctx, i)
}
