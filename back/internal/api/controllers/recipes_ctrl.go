package ctrl

import (
	dto_recipe_req "food/internal/api/http/dtos/recipe/req"
	"food/internal/api/http/mappers"
	req_Extractor "food/internal/api/http/req-extractor"
	recipe_app "food/internal/modules/recipe/application"
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
	req, status, err := req_Extractor.WorkspaceRequestExtractor[dto_recipe_req.CreateRecipeDto](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	cmd := mappers.ToRecipeCreateCmd(req.WorkspaceId, req.Dto)
	newRecipe, err := rctrl.recipeSrv.Create(r.Context(), cmd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, newRecipe, http.StatusCreated)
}

func (rctrl *RecipeCtrl) Update(w http.ResponseWriter, r *http.Request) {
	req, status, err := req_Extractor.WorkspaceRequestExtractor[dto_recipe_req.UpdateRecipeDto](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	cmd := mappers.ToRecipeUpdateCmd(req.WorkspaceId, req.Dto)
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
