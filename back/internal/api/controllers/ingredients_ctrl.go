package ctrl

import (
	"encoding/json"
	"fmt"
	dto_recipe_req "food/internal/api/http/dtos/recipe/req"
	api_validator "food/internal/api/validator"
	recipe_app "food/internal/modules/recipe/application"
	recipe_dom "food/internal/modules/recipe/domain"
	platform_http "food/internal/platform/http"
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

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found, err := ictrl.ingredientSrv.Search(r.Context(), workspaceId, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found, http.StatusOK)
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

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	found, err := ictrl.ingredientSrv.Create(r.Context(), recipe_app.CreateIngredientCmd{
		Name:        req.Name,
		BaseUnit:    req.BaseUnit,
		WorkspaceId: workspaceId,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found, http.StatusCreated)
}

func (ictrl *IngredientCtrl) Delete(w http.ResponseWriter, r *http.Request) {
	ingredientUuid, err := uuid.Parse(r.PathValue("ingredient-id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	workspaceId, ok := platform_http.GetWorkspaceIdFromContext(r.Context())
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deleteCmd := recipe_app.DeleteIngredientCmd{
		WorkspaceId: workspaceId,
		Id:          ingredientUuid,
	}

	err = ictrl.ingredientSrv.DeleteById(r.Context(), deleteCmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, fmt.Sprintf("ingredient %s deleted", ingredientUuid.String()), http.StatusAccepted)
}
