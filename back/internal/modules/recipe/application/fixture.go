package recipe_app

import (
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/google/uuid"
)

// IDs are immutable, so sharing them is safe.
var (
	saltID        = uuid.MustParse("55d03b43-82cb-46bd-8645-fa66e88bb996")
	pastaID       = uuid.MustParse("90cc7659-912a-460a-bf9e-fb04c1f9c08c")
	tomatoSauceID = uuid.MustParse("ab618a4c-e2c7-4e34-ad1f-4fe88d2e98f9")
	recipeID      = uuid.MustParse("d9578476-6f71-498b-9e49-c90bab2f6004")

	workspaceID = uuid.MustParse("b1b9f832-b233-48db-8738-fb1908ba7a91")
)

const (
	recipeName    = "salt pasta"
	recipeNewName = "tomato pasta"
)

// Ingredients are pointers (mutable), so every test gets its own copy.
func saltIngredient() *recipe_dom.Ingredient {
	return recipe_dom.NewIngredient(saltID, "salt", recipe_dom.Gram, workspaceID)
}
func pastaIngredient() *recipe_dom.Ingredient {
	return recipe_dom.NewIngredient(pastaID, "pasta", recipe_dom.Kilo, workspaceID)
}
func tomatoSauceIngredient() *recipe_dom.Ingredient {
	return recipe_dom.NewIngredient(tomatoSauceID, "tomato sauce", recipe_dom.Kilo, workspaceID)
}

func allIngredients() []*recipe_dom.Ingredient {
	return []*recipe_dom.Ingredient{saltIngredient(), pastaIngredient(), tomatoSauceIngredient()}
}

// -- requests --

func createSaltPastaReq() dto_recipe_req.CreateRecipeDto {
	return dto_recipe_req.CreateRecipeDto{
		Name: recipeName,
		Ingredients: []dto_recipe_req.IngredientBasicInfo{
			{saltID, 5, recipe_dom.Gram},
			{pastaID, 500, recipe_dom.Gram},
		},
	}
}

// updateReq builds an update of recipeID with the new name and the given lines.
func updateReq(lines ...dto_recipe_req.IngredientBasicInfo) dto_recipe_req.UpdateRecipeDto {
	return dto_recipe_req.UpdateRecipeDto{
		Id: recipeID,
		CreateRecipeDto: dto_recipe_req.CreateRecipeDto{
			Name:        recipeNewName,
			Ingredients: lines,
		},
	}
}

// -- expected domain models (explicit, no mutation of other fixtures) --

func expectedSaltPasta() *recipe_dom.Recipe {
	r := recipe_dom.NewRecipe(recipeID, recipeName)
	r.AddIngredient(saltIngredient(), recipe_dom.MustNewQuantity(5, recipe_dom.Gram))
	r.AddIngredient(pastaIngredient(), recipe_dom.MustNewQuantity(500, recipe_dom.Gram))
	return r
}

func expectedTomatoPasta() *recipe_dom.Recipe {
	r := recipe_dom.NewRecipe(recipeID, recipeNewName)
	r.AddIngredient(pastaIngredient(), recipe_dom.MustNewQuantity(500, recipe_dom.Gram))
	r.AddIngredient(tomatoSauceIngredient(), recipe_dom.MustNewQuantity(100, recipe_dom.Gram))
	return r
}
