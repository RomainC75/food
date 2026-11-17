package mappers

import (
	dto_recipe_req "food/internal/api/http/dtos/recipe/req"
	recipe_app "food/internal/modules/recipe/application"

	"github.com/google/uuid"
)

func ToRecipeCmd(workspaceId uuid.UUID, req dto_recipe_req.CreateRecipeDto) recipe_app.CreateRecipeCmd {
	ingredientBIs := make([]recipe_app.IngredientBasicInfo, 0, len(req.Ingredients))
	for _, ibi := range req.Ingredients {
		ingredientBI := recipe_app.IngredientBasicInfo{
			Id:   ibi.Id,
			Qty:  ibi.Qty,
			Unit: ibi.Unit,
		}
		ingredientBIs = append(ingredientBIs, ingredientBI)
	}
	return recipe_app.CreateRecipeCmd{
		WorkspaceId: workspaceId,
		Name:        req.Name,
		Ingredients: ingredientBIs,
	}
}
