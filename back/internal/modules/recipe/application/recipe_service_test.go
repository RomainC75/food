package recipe_app

import (
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------- Create ----------

func TestCreateRecipe(t *testing.T) {
	t.Run("should create recipe", func(t *testing.T) {
		td := NewTestDriver(t).SetNextUuids(recipeID).GivenIngredients(saltIngredient(), pastaIngredient())

		created, err := td.CreateRecipe(createSaltPastaReq())

		require.NoError(t, err)
		assert.Equal(t, expectedSaltPasta(), created)
		assert.Len(t, td.SavedRecipes(), 1)
		assert.Equal(t, expectedSaltPasta(), td.SavedRecipes()[0])
	})
	t.Run("should no create duplicated recipe", func(t *testing.T) {
		td := NewTestDriver(t).SetNextUuids(uuid.New(), recipeID).GivenIngredients(saltIngredient(), pastaIngredient())
		cmd := createSaltPastaReq()
		td.CreateRecipe(cmd)

		_, err := td.CreateRecipe(cmd)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), ErrCouldNotDuplicateRecipe.Error())
		assert.Contains(t, err.Error(), cmd.Name)
	})
}

// ---------- Update ----------

func TestUpdateRecipe_Success(t *testing.T) {
	t.Run("should update", func(t *testing.T) {
		td := NewTestDriver(t).GivenIngredients(allIngredients()...)
		td.GivenRecipe(recipeID, createSaltPastaReq())

		updated, err := td.UpdateRecipe(updateReq(
			workspaceID,
			IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
			IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
		))

		require.NoError(t, err)
		assert.Equal(t, expectedTomatoPasta(), updated)
		assert.Equal(t, []*recipe_dom.Recipe{expectedTomatoPasta()}, td.SavedRecipes())
	})
}

func TestUpdateRecipe_Errors(t *testing.T) {
	tests := []struct {
		name         string
		ingredients  []*recipe_dom.Ingredient // known by the ingredient repo
		recipeExists bool
		repoErrors   recipe_infra.RecipeRepoErrorsToSet
		cmd          UpdateRecipeCmd
		wantErr      []string // every string must be contained in err.Error()
	}{
		{
			name:         "recipe not found",
			recipeExists: false,
			cmd:          updateReq(workspaceID),
			wantErr:      []string{recipe_dom.ErrRecipeNotFound.Error()},
		},
		{
			name:         "ingredient not found",
			ingredients:  []*recipe_dom.Ingredient{saltIngredient(), pastaIngredient()}, // no tomato sauce
			recipeExists: true,
			cmd: updateReq(
				workspaceID,
				IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
				IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
			),
			wantErr: []string{ErrAddingIngredients, ErrIngredientsNotFound, tomatoSauceID.String()},
		},
		{
			name:         "wrong unit",
			ingredients:  allIngredients(),
			recipeExists: true,
			cmd: updateReq(
				workspaceID,
				IngredientBasicInfo{pastaID, 500, recipe_dom.Liter},
				IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
			),
			wantErr: []string{ErrAddingIngredients, ErrWrongUnit, pastaID.String()},
		},
		{
			name:         "negative quantity",
			ingredients:  allIngredients(),
			recipeExists: true,
			cmd: updateReq(
				workspaceID,
				IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
				IngredientBasicInfo{tomatoSauceID, -10, recipe_dom.Gram},
			),
			wantErr: []string{ErrAddingIngredients, ErrWrongQuantity, tomatoSauceID.String()},
		},
		{
			name:         "repository fails to save",
			ingredients:  allIngredients(),
			recipeExists: true,
			repoErrors:   recipe_infra.RecipeRepoErrorsToSet{UpdateError: true},
			cmd: updateReq(
				workspaceID,
				IngredientBasicInfo{pastaID, 500, recipe_dom.Gram},
				IngredientBasicInfo{tomatoSauceID, 100, recipe_dom.Gram},
			),
			wantErr: []string{recipe_dom.ErrCouldNotUpdateRecipe.Error()},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			td := NewTestDriver(t).GivenIngredients(tc.ingredients...)
			if tc.recipeExists {
				td.GivenRecipe(recipeID, createSaltPastaReq())
			}
			td.GivenRepoErrors(tc.repoErrors)

			_, err := td.UpdateRecipe(tc.cmd)

			assertErrContainsAll(t, err, tc.wantErr...)
		})
	}
}

// ---------- Delete ----------

func TestDeleteRecipe(t *testing.T) {
	t.Run("should return an error if the recipe is not found", func(t *testing.T) {
		td := NewTestDriver(t).GivenIngredients(saltIngredient(), pastaIngredient())
		td.GivenRecipe(recipeID, createSaltPastaReq())

		recipeToDelete := uuid.New()
		err := td.DeleteRecipe(workspaceID, recipeToDelete)

		assert.Contains(t, err.Error(), recipe_dom.ErrRecipeNotFound.Error())
		assert.Contains(t, err.Error(), recipeToDelete.String())
		assert.Len(t, td.SavedRecipes(), 1)
	})

	t.Run("existing recipe is deleted", func(t *testing.T) {
		td := NewTestDriver(t).GivenIngredients(saltIngredient(), pastaIngredient())
		td.GivenRecipe(recipeID, createSaltPastaReq())

		err := td.DeleteRecipe(workspaceID, recipeID)

		assert.NoError(t, err)
		assert.Empty(t, td.SavedRecipes())
	})
}
