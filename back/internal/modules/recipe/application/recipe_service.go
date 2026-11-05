package recipe_app

import (
	"context"
	"errors"
	"fmt"
	dto_recipe_req "food/internal/api/dtos/recipe/req"
	recipe_dom "food/internal/modules/recipe/domain"
	shared_domain "food/internal/shared/domain"
	"food/utils/displayer"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrRecipeNotFound    = "recipe not found"
	ErrAddingIngredients = "error adding ingredients"

	ErrIngredientsNotFound = "ingredients not found"
	ErrWrongUnit           = "wrong unit"
	ErrWrongQuantity       = "wrong quantity"
)

type ErrIngredientNotFound struct {
	IngredientNotFound []dto_recipe_req.IngredientBasicInfo
	WrongUnit          []dto_recipe_req.IngredientBasicInfo
	WrongQuantity      []dto_recipe_req.IngredientBasicInfo
}

func (einf *ErrIngredientNotFound) Error() string {
	fullStr := &strings.Builder{}
	displayer.PrettyDisplay("ERR FULL : ", einf)
	if len(einf.IngredientNotFound) > 0 {
		compileErrorMessage(einf.IngredientNotFound, fullStr, ErrIngredientsNotFound)
	}

	if len(einf.WrongUnit) > 0 {
		compileErrorMessage(einf.WrongUnit, fullStr, ErrWrongUnit)
	}
	if len(einf.WrongQuantity) > 0 {
		compileErrorMessage(einf.WrongQuantity, fullStr, ErrWrongQuantity)
	}
	return fmt.Sprintf("%s : %s", ErrAddingIngredients, fullStr.String())
}

func compileErrorMessage(slc []dto_recipe_req.IngredientBasicInfo, b *strings.Builder, errMessage string) {
	ids := make([]string, 0, len(slc))
	for _, bI := range slc {
		ids = append(ids, bI.Id.String())
	}
	nfStrList := fmt.Sprintf("%s : %s", errMessage, strings.Join(ids, " , "))
	b.Write([]byte(nfStrList + " | "))
}

func (einf *ErrIngredientNotFound) IsError() bool {
	return len(einf.IngredientNotFound) > 0 || len(einf.WrongUnit) > 0 || len(einf.WrongQuantity) > 0
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
			continue
		}
		qty, err := recipe_dom.NewQuantity(bi.Qty, bi.Unit)
		if err != nil {
			errINF.WrongUnit = append(errINF.WrongUnit, bi)
			continue
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
	errINF := &ErrIngredientNotFound{
		IngredientNotFound: make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
		WrongUnit:          make([]dto_recipe_req.IngredientBasicInfo, 0, len(r.Ingredients)),
	}
	// get
	foundRecipe, err := rSrv.recipes.GetById(ctx, r.Id)
	if err != nil {
		return nil, errors.New(ErrRecipeNotFound)
	}
	foundRecipe.SetName(r.Name)
	foundRecipe.DeleteAllIngredients()

	for _, bi := range r.Ingredients {
		foundIngredient, err := rSrv.ingredients.GetById(ctx, bi.Id)
		displayer.PrettyDisplay("get this ingredient by id: ", bi)
		if err != nil {
			errINF.IngredientNotFound = append(errINF.IngredientNotFound, bi)
			continue
		}
		qty, err := recipe_dom.NewQuantity(bi.Qty, bi.Unit)
		if err != nil {
			errINF.WrongQuantity = append(errINF.WrongQuantity, bi)
			continue
		}
		err = foundRecipe.AddIngredient(foundIngredient, qty)
		if err != nil {
			errINF.WrongUnit = append(errINF.WrongUnit, bi)
		}
	}
	if errINF.IsError() {
		return nil, errINF
	}
	err = rSrv.recipes.Update(ctx, foundRecipe)
	return foundRecipe, err
}

func (rSrv *RecipeSrv) Delete(ctx context.Context, id uuid.UUID) error {
	return rSrv.recipes.Delete(ctx, id)
}
