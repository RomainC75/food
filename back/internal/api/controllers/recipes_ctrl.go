package ctrl

import (
	"encoding/json"
	dto_recipe_req "food/internal/api/http/dtos/recipe/req"
	"food/internal/api/http/mappers"
	api_validator "food/internal/api/validator"
	recipe_app "food/internal/modules/recipe/application"
	platform_http "food/internal/platform/http"
	json_utils "food/utils/json"
	"net/http"
)

type RecipeCtrl struct {
	recipeSrv *recipe_app.RecipeSrv
}

func NewRecipeCtrl(recipeSrv *recipe_app.RecipeSrv) *RecipeCtrl {
	return &RecipeCtrl{
		recipeSrv: recipeSrv,
	}
}

func (rctrl *RecipeCtrl) Create(w http.ResponseWriter, r *http.Request) {
	var req dto_recipe_req.CreateRecipeDto
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = api_validator.Get().Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := mappers.ToRecipeCreateCmd(workspaceId, req)
	newRecipe, err := rctrl.recipeSrv.Create(r.Context(), cmd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, newRecipe, http.StatusCreated)
}

func (rctrl *RecipeCtrl) Update(w http.ResponseWriter, r *http.Request) {
	var req dto_recipe_req.UpdateRecipeDto
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = api_validator.Get().Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := mappers.ToRecipeUpdateCmd(workspaceId, req)
	newRecipe, err := rctrl.recipeSrv.Update(r.Context(), cmd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, newRecipe, http.StatusCreated)
}

// func (rctrl *RecipeCtrl) Delete(w http.ResponseWriter, r *http.Request) {
// 	ingredientUuid, err := uuid.Parse(r.PathValue("ingredient-id"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
