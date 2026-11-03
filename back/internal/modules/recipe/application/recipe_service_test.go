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
	saltId  = uuid.MustParse("55d03b43-82cb-46bd-8645-fa66e88bb996")
	saltI   = recipe_dom.NewIngredient(saltId, "salt", recipe_dom.Gram)
	pastaid = uuid.MustParse("90cc7659-912a-460a-bf9e-fb04c1f9c08c")
	pastaI  = recipe_dom.NewIngredient(pastaid, "pasta", recipe_dom.Kilo)

	recipeId    = uuid.MustParse("d9578476-6f71-498b-9e49-c90bab2f6004")
	receipeName = "salt pasta"

	receipeReq = dto_recipe_req.CreateReceipeDto{
		Name: receipeName,
		Ingredients: []dto_recipe_req.IngredientBasicInfo{
			{saltId, 5, recipe_dom.Gram},
			{pastaid, 0.5, recipe_dom.Kilo},
		},
	}
)

func TestRecipeSrv(t *testing.T) {
	recipe := recipe_dom.NewReceipe(recipeId, "salt pastas")

	salt := recipe_dom.NewIngredient(saltId, "salt", recipe_dom.Gram)
	pasta := recipe_dom.NewIngredient(pastaid, "pasta", recipe_dom.Kilo)
	recipe.AddIngredient(pasta)
	recipe.AddIngredient(salt)

	td := NewTestDriver()
	td.SetIngredientInMemory(saltI)
	td.SetIngredientInMemory(pastaI)

	td.SetNextUuid(recipeId)

	createdReceipe, err := td.CreateReceipe(receipeReq)

	expectedReceipe := recipe_dom.NewReceipe(recipeId, receipeName)
	expectedReceipe.AddIngredient(saltI)
	expectedReceipe.AddIngredient(pastaI)

	if err != nil {
		log.Println("erro : ", err.Error())
	}
	assert.Nil(t, err)
	assert.Equal(t, *expectedReceipe, *createdReceipe)
}
