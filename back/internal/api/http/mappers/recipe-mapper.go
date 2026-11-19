package mappers

import (
	dto_recipe_req "food/internal/api/http/dtos/recipe/req"
	recipe_app "food/internal/modules/recipe/application"

	"github.com/google/uuid"
)

func ToRecipeCreateCmd(workspaceId uuid.UUID, req dto_recipe_req.CreateRecipeDto) recipe_app.CreateRecipeCmd {
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

func ToRecipeUpdateCmd(workspaceId uuid.UUID, req dto_recipe_req.UpdateRecipeDto) recipe_app.UpdateRecipeCmd {
	ingredientBIs := make([]recipe_app.IngredientBasicInfo, 0, len(req.Ingredients))
	for _, ibi := range req.Ingredients {
		ingredientBI := recipe_app.IngredientBasicInfo{
			Id:   ibi.Id,
			Qty:  ibi.Qty,
			Unit: ibi.Unit,
		}
		ingredientBIs = append(ingredientBIs, ingredientBI)
	}
	return recipe_app.UpdateRecipeCmd{
		WorkspaceId: workspaceId,
		Id:          req.Id,
		Name:        req.Name,
		Ingredients: ingredientBIs,
	}
}
