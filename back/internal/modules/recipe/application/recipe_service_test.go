package recipe_app

import (
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	saltId         = uuid.MustParse("55d03b43-82cb-46bd-8645-fa66e88bb996")
	saltI          = recipe_dom.NewIngredient(saltId, "salt", recipe_dom.Gram)
	saltQty        = recipe_dom.MustNewQuantity(5, recipe_dom.Gram)
	pastaid        = uuid.MustParse("90cc7659-912a-460a-bf9e-fb04c1f9c08c")
	pastaI         = recipe_dom.NewIngredient(pastaid, "pasta", recipe_dom.Kilo)
	pastaQty       = recipe_dom.MustNewQuantity(500, recipe_dom.Gram)
	sauceTomateid  = uuid.MustParse("90cc7659-912a-460a-bf9e-fb04c1f9c08c")
	sauceTomateI   = recipe_dom.NewIngredient(sauceTomateid, "sauce omate", recipe_dom.Kilo)
	sauceTomateQty = recipe_dom.MustNewQuantity(100, recipe_dom.Gram)

	recipeId   = uuid.MustParse("d9578476-6f71-498b-9e49-c90bab2f6004")
	recipeName = "salt pasta"

	recipeReq = dto_recipe_req.CreateRecipeDto{
		Name: recipeName,
		Ingredients: []dto_recipe_req.IngredientBasicInfo{
			{saltId, 5, recipe_dom.Gram},
			{pastaid, 500, recipe_dom.Gram},
		},
	}

	updateRecipeReq = dto_recipe_req.UpdateRecipeDto{
		Id: recipeId,
		CreateRecipeDto: dto_recipe_req.CreateRecipeDto{
			Name: recipeName,
			Ingredients: []dto_recipe_req.IngredientBasicInfo{
				{pastaid, 500, recipe_dom.Gram},
				{sauceTomateid, 100, recipe_dom.Gram},
			},
		},
	}
)

func TestRecipeSrv(t *testing.T) {
	td := NewTestDriver()
	td.SetIngredientInMemory(saltI)
	td.SetIngredientInMemory(pastaI)
	td.SetIngredientInMemory((sauceTomateI))
	td.SetNextUuid(recipeId)

	createdRecipe, err := td.CreateRecipe(recipeReq)

	// -- create expected model to compare --
	expectedRecipe := recipe_dom.NewRecipe(recipeId, recipeName)
	expectedRecipe.AddIngredient(saltI, saltQty)
	expectedRecipe.AddIngredient(pastaI, pastaQty)

	if err != nil {
		log.Println("erro : ", err.Error())
	}
	assert.Nil(t, err)
	assert.Equal(t, *expectedRecipe, *createdRecipe)

	// update
	updatedRecipe, err := td.UpdateRecipe(updateRecipeReq)
	// expected
	expectedRecipe.SetName(updateRecipeReq.Name)
	expectedRecipe.DeleteIngredient(saltId)
	expectedRecipe.AddIngredient(sauceTomateI, sauceTomateQty)

	assert.Equal(t, expectedRecipe, updatedRecipe)
	savedRecipes := td.GetSavedRecipes()
	assert.Len(t, savedRecipes, 1)
	assert.Equal(t, savedRecipes[0], expectedRecipe)

	// delete
	err = td.DeleteRecipe(updatedRecipe.Id())
	assert.Nil(t, err)
	savedRecipes = td.GetSavedRecipes()
	assert.Len(t, savedRecipes, 0)
}

func TestDeleteRecipe(t *testing.T) {
	td := NewTestDriver()
	createdRecipe, err := td.SetupRecipe(recipeReq, []*recipe_dom.Ingredient{saltI, pastaI})
	assert.Nil(t, err)

	err = td.DeleteRecipe(uuid.New())
	assert.EqualError(t, err, ErrRecipeNotFound)
	assert.Len(t, td.GetSavedRecipes(), 1)

	err = td.DeleteRecipe(createdRecipe.Id())
	assert.Nil(t, err)
	assert.Len(t, td.GetSavedRecipes(), 0)

}
