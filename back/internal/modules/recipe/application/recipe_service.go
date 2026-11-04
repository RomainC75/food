package recipe_app

import (
	"context"
	"errors"
	"fmt"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	shared_domain "food/internal/shared/domain"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrRecipeNotFound = "recipe not found"
)

type ErrIngredientNotFound struct {
	IngredientNotFound []dto_recipe_req.IngredientBasicInfo
	WrongUnit          []dto_recipe_req.IngredientBasicInfo
}

func (einf *ErrIngredientNotFound) Error() string {
	wIds := make([]string, 0, len(einf.IngredientNotFound))
	var nfStrList string
	fullStr := &strings.Builder{}

	if len(einf.IngredientNotFound) > 0 {

		for _, bI := range einf.IngredientNotFound {
			wIds = append(wIds, bI.Id.String())
		}
		nfStrList = fmt.Sprintf("ingredients not found : %s", strings.Join(wIds, " , "))
		fullStr.Write([]byte(nfStrList + " | "))
	}

	wuIds := make([]string, 0, len(einf.WrongUnit))
	var wuStrList string
	if len(einf.IngredientNotFound) > 0 {

		for _, bI := range einf.IngredientNotFound {
			wuIds = append(wuIds, bI.Id.String())
		}
		wuStrList = fmt.Sprintf("wrongQty for : %s", strings.Join(wuIds, " , "))
		fullStr.Write([]byte(wuStrList))
	}

	return fmt.Sprintf("error adding ingredients : %s ", fullStr.String())
}

type RecipeSrv struct {
	uuidGen     shared_domain.Uuids
	ingredients recipe_dom.Ingredients
	recipes     recipe_dom.Recipes
}

func NewRecipeSrv(uuidGen shared_domain.Uuids, i recipe_dom.Ingredients, r recipe_dom.Recipes) *RecipeSrv {
	return &RecipeSrv{
		uuidGen:     uuidGen,
		ingredients: i,
		recipes:     r,
	}
}

func (rSrv *RecipeSrv) Create(ctx context.Context, r dto_recipe_req.CreateRecipeDto) (*recipe_dom.Recipe, error) {
	errINF := ErrIngredientNotFound{
		IngredientNotFound: make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
		WrongUnit:          make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
	}

	recipe := recipe_dom.NewRecipe(rSrv.uuidGen.Generate(), r.Name)

	for _, bi := range r.Ingredients {
		i, err := rSrv.ingredients.GetById(ctx, bi.Id)
		if err != nil {
			errINF.IngredientNotFound = append(errINF.IngredientNotFound, bi)
		}
		qty, err := recipe_dom.NewQuantity(bi.Qty, bi.Unit)
		if err != nil {
			errINF.WrongUnit = append(errINF.WrongUnit, bi)
		}
		recipe.AddIngredient(i, qty)
	}

	if len(errINF.IngredientNotFound) != 0 || len(errINF.WrongUnit) != 0 {
		return nil, &errINF
	}

	err := rSrv.recipes.Create(ctx, recipe)
	return recipe, err
}

func (rSrv *RecipeSrv) Update(ctx context.Context, r dto_recipe_req.UpdateRecipeDto) (*recipe_dom.Recipe, error) {
	errINF := ErrIngredientNotFound{
		IngredientNotFound: make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
		WrongUnit:          make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
	}
	// get
	foundRecipe, err := rSrv.recipes.GetById(ctx, r.Id)
	if err != nil {
		return nil, errors.New(ErrRecipeNotFound)
	}
	foundRecipe.SetName(r.Name)

	newRecipeIngredients := make([]recipe_dom.RecipeIngredient, 0, len(r.Ingredients))
	for _, bi := range r.Ingredients {
		_, err := rSrv.ingredients.GetById(ctx, bi.Id)
		if err != nil {
			errINF.IngredientNotFound = append(errINF.IngredientNotFound, bi)
		}
		qty, err := recipe_dom.NewQuantity(bi.Qty, bi.Unit)
		if err != nil {
			errINF.WrongUnit = append(errINF.WrongUnit, bi)
		}
		newRecipeIngredients = append(newRecipeIngredients, recipe_dom.RecipeIngredient{
			IngredientID: bi.Id,
			Quantity:     qty,
		})
	}

	foundRecipe.ReplaceIngredients(newRecipeIngredients)
	err = rSrv.recipes.Update(ctx, foundRecipe)
	return foundRecipe, err
}

func (rSrv *RecipeSrv) Delete(ctx context.Context, id uuid.UUID) error {
	return rSrv.recipes.Delete(ctx, id)
}
