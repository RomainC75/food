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

	ingredientSalt = dto_recipe_req.CreateIngredientDto{
		Name:     saltName,
		BaseUnit: saltBaseUnit,
	}

	ingredientTomato = dto_recipe_req.CreateIngredientDto{
		Name:     tomatoName,
		BaseUnit: tomatoBaseUnit,
	}

	userUuid      = uuid.MustParse("89a9525a-0e49-4e1b-9990-a5042ca19c00")
	workspaceUuid = uuid.MustParse("b1b9f832-b233-48db-8738-fb1908ba7a91")

	expectedSavedIngredient = recipe_dom.NewIngredient(saltUuid, saltName, recipe_dom.BaseUnit(saltBaseUnit), workspaceUuid)
)

func TestCreateIngredient(t *testing.T) {
	t.Run("should create a ingredient", func(t *testing.T) {
		td := NewTestDriver(t).SetContext(userUuid, workspaceUuid)
		td.SetNextUuid(saltUuid)
		createdIngredient, err := td.CreateIngredient(workspaceUuid, ingredientSalt)

		assert.Nil(t, err)
		assert.Equal(t, expectedSavedIngredient, createdIngredient)

		savedIngredients := td.SavedIngredients()
		assert.Len(t, savedIngredients, 1)
		assert.Equal(t, expectedSavedIngredient, savedIngredients[0])
	})
	t.Run("should not create a duplicate ingredient", func(t *testing.T) {
		td := NewTestDriver(t).SetContext(userUuid, workspaceUuid)
		td.SetNextUuid(saltUuid)
		td.CreateIngredient(workspaceUuid, ingredientSalt)
		_, err := td.CreateIngredient(workspaceUuid, ingredientSalt)

		assert.EqualError(t, err, ErrIngredientAlreadyExists)
	})
}

func TestGetIngredientById(t *testing.T) {
	td := NewTestDriver(t).SetContext(userUuid, workspaceUuid)
	td.SetNextUuids(saltUuid, tomatoUuid)
	td.CreateIngredient(workspaceUuid, ingredientSalt)
	td.CreateIngredient(workspaceUuid, ingredientTomato)

	searchFilter := recipe_dom.IngredientFilter{
		Name: tomatoName,
	}

	found, err := td.Search(workspaceUuid, searchFilter)
	assert.Nil(t, err)
	assert.Equal(t, td.SavedIngredients()[1], found)
}

func TestSearchIngredient(t *testing.T) {
	td := NewTestDriver(t).SetContext(userUuid, workspaceUuid)
	td.SetNextUuids(saltUuid, tomatoUuid)
	td.CreateIngredient(workspaceUuid, ingredientSalt)
	td.CreateIngredient(workspaceUuid, ingredientTomato)

	searchFilter := recipe_dom.IngredientFilter{
		Name: tomatoName,
	}
	found, err := td.Search(workspaceUuid, searchFilter)
	assert.Nil(t, err)
	assert.Equal(t, td.SavedIngredients()[1], found)
}

func TestUpdateIngredient(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(recipe_dom.NewIngredient(saltUuid, saltName, saltBaseUnit, workspaceUuid))
	falseIdingredient := dto_recipe_req.UpdateIngredientDto{
		Id: saltUuid,
		CreateIngredientDto: dto_recipe_req.CreateIngredientDto{
			Name:     tomatoName,
			BaseUnit: tomatoBaseUnit,
		},
	}

	t.Run("cannot get the ingredient", func(t *testing.T) {
		td.GivenIngredientRepoErrors(true, false)
		_, err := td.UpdateIngredient(workspaceUuid, falseIdingredient)
		assert.EqualError(t, err, recipe_dom.ErrIngredientNotFound)

	})

	t.Run("should not update ingredient", func(t *testing.T) {
		td.GivenIngredientRepoErrors(false, true)
		_, err := td.UpdateIngredient(workspaceUuid, falseIdingredient)
		assert.EqualError(t, err, recipe_dom.ErrIngredientCouldNotBeUpdated)
	})

	t.Run("should update ingredient", func(t *testing.T) {
		td.GivenIngredientRepoErrors(false, false)
		updatedIngredient, err := td.UpdateIngredient(workspaceUuid, falseIdingredient)

		expectedIngredient := recipe_dom.NewIngredient(saltUuid, tomatoName, tomatoBaseUnit, workspaceUuid)

		assert.Nil(t, err)
		assert.Equal(t, expectedIngredient, updatedIngredient)

		assert.Len(t, td.SavedIngredients(), 1)
		assert.Equal(t, expectedIngredient, td.SavedIngredients()[0])
	})

}

func TestDeleteIngredient(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(expectedSavedIngredient)
	err := td.DeleteIngredient(workspaceUuid, saltUuid)

	assert.Nil(t, err)
	deletedIngredient := td.GetDeletedLastIngredient()

	assert.Equal(t, expectedSavedIngredient, deletedIngredient)

}
