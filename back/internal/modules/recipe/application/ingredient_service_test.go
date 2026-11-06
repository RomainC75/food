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

	tomatoUuid     = uuid.MustParse("ad3635f2-d87c-4f7c-83a6-9b8846eaaba2")
	tomatoName     = "tomato"
	tomatoBaseUnit = recipe_dom.Piece

	ingredient1 = dto_recipe_req.CreateIngredientDto{
		Name:     saltName,
		BaseUnit: saltBaseUnit,
	}

	ingredient2 = dto_recipe_req.CreateIngredientDto{
		Name:     tomatoName,
		BaseUnit: tomatoBaseUnit,
	}

	expectedSavedIngredient = recipe_dom.NewIngredient(saltUuid, saltName, recipe_dom.BaseUnit(saltBaseUnit))
)

func TestCreateIngredient(t *testing.T) {
	td := NewTestDriver(t)
	td.SetNextUuid(saltUuid)
	createdIngredient, err := td.CreateIngredient(ingredient1)

	assert.Nil(t, err)
	assert.Equal(t, expectedSavedIngredient, createdIngredient)

	savedIngredients := td.SavedIngredients()
	assert.Len(t, savedIngredients, 1)
	assert.Equal(t, expectedSavedIngredient, savedIngredients[0])
}

func TestSearchIngredient(t *testing.T) {
	td := NewTestDriver(t)
	td.SetNextUuid(saltUuid)
	td.CreateIngredient(ingredient1)
	td.CreateIngredient(ingredient2)

	searchFilter := recipe_dom.IngredientFilter{
		Name: tomatoName,
	}
	found, err := td.Search(searchFilter)
	assert.Nil(t, err)
	assert.Equal(t, td.SavedIngredients()[1], found)
}

func TestUpdateIngredient(t *testing.T) {

	td := NewTestDriver(t).GivenIngredients(recipe_dom.NewIngredient(saltUuid, saltName, saltBaseUnit))

	// ingNot found
	falseIdingredient := dto_recipe_req.UpdateIngredientDto{
		Id: tomatoUuid,
		CreateIngredientDto: dto_recipe_req.CreateIngredientDto{
			Name:     tomatoName,
			BaseUnit: tomatoBaseUnit,
		},
	}
	_, err := td.UpdateIngredient(falseIdingredient)
	assert.EqualError(t, err, recipe_dom.ErrIngredientNotFound)

	// ? ing ok, but cannot update

	// ? everything ok
}

func TestDeleteIngredient(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(expectedSavedIngredient)
	err := td.DeleteIngredient(saltUuid)

	assert.Nil(t, err)
	deletedIngredient := td.GetDeletedLastIngredient()

	assert.Equal(t, expectedSavedIngredient, deletedIngredient)

}
