package ctrl

import (
	"context"
	"encoding/json"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	api_validator "food/internal/api/validator"
	recipe_app "food/internal/modules/recipe/application"
	recipe_dom "food/internal/modules/recipe/domain"
	recipe_infra "food/internal/modules/recipe/infra"
	recipe_infra_init "food/internal/modules/recipe/infra/init"
	shared_infra "food/internal/shared/infra"
	json_utils "food/utils/json"
	"net/http"
)

type IngredientCtrl struct {
	ingredientSrv *recipe_app.IngredientSrv
}

func NewIngredientCtrl() *IngredientCtrl {
	return &IngredientCtrl{
		ingredientSrv: recipe_app.NewIngredientSrv(
			shared_infra.NewInMemUuidGen(),
			recipe_infra.NewInMemoryIngredients().SetIngredientsInMem(recipe_infra_init.Ingredients...),
		),
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

	ctx := context.Background()
	found, err := ictrl.ingredientSrv.Create(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json_utils.JsonResponse(w, found)
}
