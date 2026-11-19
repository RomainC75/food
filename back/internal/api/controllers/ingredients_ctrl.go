package ctrl

import (
	"fmt"
	dto_recipe_req "food/internal/api/http/dtos/recipe/req"
	req_extractor "food/internal/api/http/req-extractor"
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
	req, status, err := req_extractor.WorkspaceRequestExtractor[recipe_dom.IngredientFilter](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	found, err := ictrl.ingredientSrv.Search(r.Context(), req.WorkspaceId, req.Dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found, http.StatusOK)
}

func (ictrl *IngredientCtrl) CreateIncredient(w http.ResponseWriter, r *http.Request) {
	req, status, err := req_extractor.WorkspaceRequestExtractor[dto_recipe_req.CreateIngredientDto](r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	found, err := ictrl.ingredientSrv.Create(r.Context(), recipe_app.CreateIngredientCmd{
		Name:        req.Dto.Name,
		BaseUnit:    req.Dto.BaseUnit,
		WorkspaceId: req.WorkspaceId,
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
