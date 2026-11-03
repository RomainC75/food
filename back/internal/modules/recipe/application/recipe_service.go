package recipe_app

import (
	"context"
	"fmt"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	shared_domain "food/internal/shared/domain"
	"strings"
)

type ErrIngredientNotFound struct {
	IngredientBasicInfo []dto_recipe_req.IngredientBasicInfo
}

func (einf *ErrIngredientNotFound) Error() string {
	ids := strings.Builder{}
	for _, bI := range einf.IngredientBasicInfo {
		ids.Write([]byte(bI.Id.String() + " - "))
	}
	return fmt.Sprintf("ingredients not found : %s", ids.String())
}

type RecipeSrv struct {
	uuidGen     shared_domain.Uuids
	ingredients recipe_dom.Ingredients
	receipes    recipe_dom.Receipes
}

func NewRecipeSrv(uuidGen shared_domain.Uuids, i recipe_dom.Ingredients, r recipe_dom.Receipes) *RecipeSrv {
	return &RecipeSrv{
		uuidGen:     uuidGen,
		ingredients: i,
		receipes:    r,
	}
}

func (rSrv *RecipeSrv) Create(ctx context.Context, r dto_recipe_req.CreateReceipeDto) (*recipe_dom.Receipe, error) {
	errINF := ErrIngredientNotFound{
		IngredientBasicInfo: make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
	}

	receipe := recipe_dom.NewReceipe(rSrv.uuidGen.Generate(), r.Name)

	for _, bi := range r.Ingredients {
		i, err := rSrv.ingredients.GetById(ctx, bi.Id)
		if err != nil {
			errINF.IngredientBasicInfo = append(errINF.IngredientBasicInfo, bi)
		}
		receipe.AddIngredient(i)
	}

	if len(errINF.IngredientBasicInfo) != 0 {
		return nil, &errINF
	}

	err := rSrv.receipes.Save(ctx, receipe)
	return receipe, err
}
