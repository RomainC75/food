package recipe_app

import (
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	saltUuid     = uuid.MustParse("4d35f8d4-16d2-4d95-b974-fe84fd7ca601")
	saltName     = "salt"
	saltBaseUnit = recipe_dom.Gram

	expectedSavedIngredient = recipe_dom.NewIngredient(saltUuid, saltName, recipe_dom.BaseUnit(saltBaseUnit))
)

func TestIngredientSrv(t *testing.T) {
	ingredient := dto_recipe_req.CreateIngredient{
		Name:     saltName,
		BaseUnit: saltBaseUnit,
	}

	td := NewTestDriver()
	td.SetNextUuid(saltUuid)
	createdIngredient, err := td.CreateIngredient(ingredient)

	assert.Nil(t, err)
	assert.Equal(t, expectedSavedIngredient, createdIngredient)

	savedIngredients := td.GetSavedIngredients()
	assert.Len(t, savedIngredients, 1)
	assert.Equal(t, expectedSavedIngredient, savedIngredients[0])
}
