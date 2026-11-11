package ctrl

import (
	"context"
	"encoding/json"
	"fmt"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	api_validator "food/internal/api/validator"
	recipe_app "food/internal/modules/recipe/application"
	recipe_dom "food/internal/modules/recipe/domain"
	json_utils "food/utils/json"
	"net/http"

	"github.com/google/uuid"
)

type IngredientCtrl struct {
	ingredientSrv *recipe_app.IngredientSrv
}

func NewIngredientCtrl(ingredientSrv *recipe_app.IngredientSrv) *IngredientCtrl {
	return &IngredientCtrl{
		ingredientSrv: ingredientSrv,
	}
}

func (ictrl *IngredientCtrl) Search(w http.ResponseWriter, r *http.Request) {
	var req recipe_dom.IngredientFilter
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

	ctx := context.Background()
	found, err := ictrl.ingredientSrv.Search(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found)
}

func (ictrl *IngredientCtrl) CreateIncredient(w http.ResponseWriter, r *http.Request) {
	var req dto_recipe_req.CreateIngredientDto
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

	found, err := ictrl.ingredientSrv.Create(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found)
}

func (ictrl *IngredientCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	ingredientUuid, err := uuid.Parse(r.PathValue("ingredient-id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ictrl.ingredientSrv.DeleteById(r.Context(), ingredientUuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, fmt.Sprintf("ingredient %s deleted", ingredientUuid.String()))
}
