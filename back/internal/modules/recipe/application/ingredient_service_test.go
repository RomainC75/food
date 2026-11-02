package recipe_app

import (
	"context"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	shared_infra "food/internal/shared/infra"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	saltUuid     = uuid.MustParse("4d35f8d4-16d2-4d95-b974-fe84fd7ca601")
	saltName     = "salt"
	saltBaseUnit = recipe_dom.Gram
)

func TestIngredientSrv(t *testing.T) {

	fakeUuid := shared_infra.NewInMemUuidGen().AddExpected(saltUuid)
	inMemIngredients := recipe_infra.NewInMemoryIngredients()
	ingredientSrv := NewIngredientSrv(fakeUuid, inMemIngredients)

	ctx := context.Background()
	ingredient := dto_recipe_req.CreateIngredient{
		Name:     saltName,
		BaseUnit: saltBaseUnit,
	}
	createdIngredient, err := ingredientSrv.Create(ctx, ingredient)
	assert.Nil(t, err)

	expectedIngredient := recipe_dom.NewIngredient(saltUuid, saltName, recipe_dom.BaseUnit(saltBaseUnit))
	assert.Equal(t, expectedIngredient, createdIngredient)

	savedIngredients := inMemIngredients.GetSaved()
	assert.Len(t, savedIngredients, 1)
	assert.Equal(t, expectedIngredient, savedIngredients[0])
}
