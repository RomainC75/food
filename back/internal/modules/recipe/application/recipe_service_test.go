package recipe_app

import (
	"errors"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	"food/utils/displayer"
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
	sauceTomateid  = uuid.MustParse("ab618a4c-e2c7-4e34-ad1f-4fe88d2e98f9")
	sauceTomateI   = recipe_dom.NewIngredient(sauceTomateid, "sauce omate", recipe_dom.Kilo)
	sauceTomateQty = recipe_dom.MustNewQuantity(100, recipe_dom.Gram)

	recipeId      = uuid.MustParse("d9578476-6f71-498b-9e49-c90bab2f6004")
	recipeName    = "salt pasta"
	recipeNewName = "tomato pasta"

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
			Name: recipeNewName,
			Ingredients: []dto_recipe_req.IngredientBasicInfo{
				{pastaid, 500, recipe_dom.Gram},
				{sauceTomateid, 100, recipe_dom.Gram},
			},
		},
	}

	updateRecipeReqWrongUnit = dto_recipe_req.UpdateRecipeDto{
		Id: recipeId,
		CreateRecipeDto: dto_recipe_req.CreateRecipeDto{
			Name: recipeNewName,
			Ingredients: []dto_recipe_req.IngredientBasicInfo{
				{pastaid, 500, recipe_dom.Liter},
				{sauceTomateid, 100, recipe_dom.Gram},
			},
		},
	}

	updateRecipeReqNegariveQuantity = dto_recipe_req.UpdateRecipeDto{
		Id: recipeId,
		CreateRecipeDto: dto_recipe_req.CreateRecipeDto{
			Name: recipeNewName,
			Ingredients: []dto_recipe_req.IngredientBasicInfo{
				{pastaid, 500, recipe_dom.Gram},
				{sauceTomateid, -10, recipe_dom.Gram},
			},
		},
	}
)

func TestCreateRecipe(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(saltI, pastaI)

	createdRecipe := td.GivenRecipe(recipeId, recipeReq)

	// -- create expected model to compare --
	expectedRecipe := recipe_dom.NewRecipe(recipeId, recipeName)
	expectedRecipe.AddIngredient(saltI, saltQty)
	expectedRecipe.AddIngredient(pastaI, pastaQty)

	assert.Equal(t, *expectedRecipe, *createdRecipe)

}

type UpdateRecipeCase struct {
	TestName              string
	RecipeRepoErrorsToSet recipe_infra.RecipeRepoErrorsToSet
	AvailableIngredients  []*recipe_dom.Ingredient
	AvailableRecipe       dto_recipe_req.CreateRecipeDto
	UpdateReq             dto_recipe_req.UpdateRecipeDto
	ExpectedError         error
	ExpectedUpdateErrors  []string
}

var updateCases = []UpdateRecipeCase{
	{
		"should update",
		recipe_infra.RecipeRepoErrorsToSet{},
		[]*recipe_dom.Ingredient{
			saltI, pastaI, sauceTomateI,
		},
		recipeReq,
		updateRecipeReq,
		// VERIFY correctly !!!!! ????
		nil,
		[]string{},
	},
	{
		"should not find the recipe to update",
		recipe_infra.RecipeRepoErrorsToSet{},
		[]*recipe_dom.Ingredient{},
		dto_recipe_req.CreateRecipeDto{},
		updateRecipeReq,
		errors.New(ErrRecipeNotFound),
		[]string{},
	},
	{
		"should not find the ingredients for the recipe",
		recipe_infra.RecipeRepoErrorsToSet{},
		[]*recipe_dom.Ingredient{
			saltI, pastaI,
		},
		recipeReq,
		updateRecipeReq,
		errors.New(ErrRecipeNotFound),
		[]string{
			ErrAddingIngredients,
			ErrIngredientsNotFound,
			sauceTomateid.String(),
		},
	},
	{
		"should have a wrong unit in one ingredient",
		recipe_infra.RecipeRepoErrorsToSet{},
		[]*recipe_dom.Ingredient{
			saltI, pastaI, sauceTomateI,
		},
		recipeReq,
		updateRecipeReqWrongUnit,
		errors.New(ErrRecipeNotFound),
		[]string{
			ErrAddingIngredients,
			ErrWrongUnit,
			pastaid.String(),
		},
	},
	{
		"should return an error if the quantity to save is negative",
		recipe_infra.RecipeRepoErrorsToSet{},
		[]*recipe_dom.Ingredient{
			saltI, pastaI, sauceTomateI,
		},
		recipeReq,
		updateRecipeReqNegariveQuantity,
		errors.New(ErrRecipeNotFound),
		[]string{
			ErrAddingIngredients,
			ErrWrongQuantity,
			sauceTomateid.String(),
		},
	},
	{
		"should return an error because the recipe could not be saved",
		recipe_infra.RecipeRepoErrorsToSet{
			UpdateError: true,
		},
		[]*recipe_dom.Ingredient{
			saltI, pastaI, sauceTomateI,
		},
		recipeReq,
		updateRecipeReq,
		nil,
		[]string{},
	},
}

func TestUpdateRecipe(t *testing.T) {
	for _, cs := range updateCases {
		t.Run(cs.TestName, func(t *testing.T) {
			td := NewTestDriver(t).GivenIngredients(cs.AvailableIngredients...)
			td.SetRecipeRepoErrors(cs.RecipeRepoErrorsToSet)
			if cs.AvailableRecipe.Name == "" {
				_, err := td.UpdateRecipe(updateRecipeReq)
				assert.EqualError(t, err, ErrRecipeNotFound)
				return
			}

			createdRecipe := td.GivenRecipe(recipeId, recipeReq)
			displayer.PrettyDisplay("available ingredients : ", td.GetSavedIngredients())
			assert.Len(t, td.GetSavedRecipes(), 1)

			// good Id
			updatedRecipe, err := td.UpdateRecipe(cs.UpdateReq)
			displayer.PrettyDisplay("err : ", err)
			// errors trying to add ingredients
			if cs.ExpectedError != nil {
				if err == nil {
					t.Fail()
					return
				}
				td.assertErrContainsAll(t, err, cs.ExpectedUpdateErrors...)
				return
			}
			// repo error
			if cs.RecipeRepoErrorsToSet.UpdateError {
				assert.EqualError(t, err, recipe_dom.ErrCouldNotUpdateRecipe)
				return
			}
			assert.Nil(t, err)
			assert.Len(t, td.GetSavedRecipes(), 1)

			// update the createdRecipe manually
			// !! Make it programmatik !!
			createdRecipe.AddIngredient(sauceTomateI, sauceTomateQty)
			createdRecipe.DeleteIngredient(saltId)
			createdRecipe.SetName(updateRecipeReq.Name)

			assert.Equal(t, createdRecipe, updatedRecipe)
			savedRecipes := td.GetSavedRecipes()
			assert.Len(t, savedRecipes, 1)
			assert.Equal(t, savedRecipes[0], createdRecipe)
		})

	}

}

func TestDeleteRecipe(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(saltI, pastaI)
	createdRecipe := td.GivenRecipe(recipeId, recipeReq)

	err := td.DeleteRecipe(uuid.New())
	assert.EqualError(t, err, ErrRecipeNotFound)
	assert.Len(t, td.GetSavedRecipes(), 1)

	err = td.DeleteRecipe(createdRecipe.Id())
	assert.Nil(t, err)
	assert.Len(t, td.GetSavedRecipes(), 0)

}
