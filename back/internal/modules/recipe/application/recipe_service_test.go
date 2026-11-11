package recipe_app

import (
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------- Create ----------

func TestCreateRecipe(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(saltIngredient(), pastaIngredient())

	created, err := td.CreateRecipe(workspaceID, recipeID, createSaltPastaReq())

	require.NoError(t, err)
	assert.Equal(t, expectedSaltPasta(), created)
	assert.Equal(t, []*recipe_dom.Recipe{expectedSaltPasta()}, td.SavedRecipes())
}

// ---------- Update ----------

func TestUpdateRecipe_Success(t *testing.T) {
	td := NewTestDriver(t).GivenIngredients(allIngredients()...)
	td.GivenRecipe(workspaceID, recipeID, createSaltPastaReq())

	updated, err := td.UpdateRecipe(workspaceID, updateReq(
		dto_recipe_req.IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
		dto_recipe_req.IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
	))

	require.NoError(t, err)
	assert.Equal(t, expectedTomatoPasta(), updated)
	assert.Equal(t, []*recipe_dom.Recipe{expectedTomatoPasta()}, td.SavedRecipes())
}

func TestUpdateRecipe_Errors(t *testing.T) {
	tests := []struct {
		name         string
		ingredients  []*recipe_dom.Ingredient // known by the ingredient repo
		recipeExists bool
		repoErrors   recipe_infra.RecipeRepoErrorsToSet
		req          dto_recipe_req.UpdateRecipeDto
		wantErr      []string // every string must be contained in err.Error()
	}{
		{
			name:         "recipe not found",
			recipeExists: false,
			req:          updateReq(),
			wantErr:      []string{ErrRecipeNotFound},
		},
		{
			name:         "ingredient not found",
			ingredients:  []*recipe_dom.Ingredient{saltIngredient(), pastaIngredient()}, // no tomato sauce
			recipeExists: true,
			req: updateReq(
				dto_recipe_req.IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
				dto_recipe_req.IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
			),
			wantErr: []string{ErrAddingIngredients, ErrIngredientsNotFound, tomatoSauceID.String()},
		},
		{
			name:         "wrong unit",
			ingredients:  allIngredients(),
			recipeExists: true,
			req: updateReq(
				dto_recipe_req.IngredientBasicInfo{pastaID, 500, recipe_dom.Liter},
				dto_recipe_req.IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
			),
			wantErr: []string{ErrAddingIngredients, ErrWrongUnit, pastaID.String()},
		},
		{
			name:         "negative quantity",
			ingredients:  allIngredients(),
			recipeExists: true,
			req: updateReq(
				dto_recipe_req.IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
				dto_recipe_req.IngredientBasicInfo{tomatoSauceID, -10, recipe_dom.Gram},
			),
			wantErr: []string{ErrAddingIngredients, ErrWrongQuantity, tomatoSauceID.String()},
		},
		{
			name:         "repository fails to save",
			ingredients:  allIngredients(),
			recipeExists: true,
			repoErrors:   recipe_infra.RecipeRepoErrorsToSet{UpdateError: true},
			req: updateReq(
				dto_recipe_req.IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
				dto_recipe_req.IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
			),
			wantErr: []string{recipe_dom.ErrCouldNotUpdateRecipe},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			td := NewTestDriver(t).GivenIngredients(tc.ingredients...)
			if tc.recipeExists {
				td.GivenRecipe(workspaceID, recipeID, createSaltPastaReq())
			}
			td.GivenRepoErrors(tc.repoErrors)

			_, err := td.UpdateRecipe(workspaceID, tc.req)

			assertErrContainsAll(t, err, tc.wantErr...)
		})
	}
}

// ---------- Delete ----------

func TestDeleteRecipe(t *testing.T) {
	t.Run("unknown id returns not found and keeps the recipe", func(t *testing.T) {
		td := NewTestDriver(t).GivenIngredients(saltIngredient(), pastaIngredient())
		td.GivenRecipe(workspaceID, recipeID, createSaltPastaReq())

		err := td.DeleteRecipe(uuid.New())

		assert.EqualError(t, err, ErrRecipeNotFound)
		assert.Len(t, td.SavedRecipes(), 1)
	})

	t.Run("existing recipe is deleted", func(t *testing.T) {
		td := NewTestDriver(t).GivenIngredients(saltIngredient(), pastaIngredient())
		td.GivenRecipe(workspaceID, recipeID, createSaltPastaReq())

		err := td.DeleteRecipe(recipeID)

		assert.NoError(t, err)
		assert.Empty(t, td.SavedRecipes())
	})
}
