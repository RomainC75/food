package recipe_infra

import (
	"context"
	"errors"
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/google/uuid"
)

var (
	ErrIngredientNotFound = "ingredient not found"
)

type InMemoryIngredients struct {
	savedIngredients   []recipe_dom.Ingredient
	expectedIngredient recipe_dom.Ingredient
	deletedIngredient  recipe_dom.Ingredient
}

func NewInMemoryIngredients() *InMemoryIngredients {
	return &InMemoryIngredients{
		savedIngredients: make([]recipe_dom.Ingredient, 0, 10),
	}
}

func (imi *InMemoryIngredients) Save(ctx context.Context, ingredient recipe_dom.Ingredient) error {
	imi.savedIngredients = append(imi.savedIngredients, ingredient)
	return nil
}

func (imi *InMemoryIngredients) GetById(ctx context.Context, id uuid.UUID) (recipe_dom.Ingredient, error) {
	for _, i := range imi.savedIngredients {
		if i.Id() == id {
			return i, nil
		}
	}
	return recipe_dom.Ingredient{}, errors.New(ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) Search(ctx context.Context, f recipe_dom.IngredientFilter) (recipe_dom.Ingredient, error) {
	for _, i := range imi.savedIngredients {
		if f.Name != "" && i.Name() != f.Name {
			continue
		}
		if f.BaseUnit != "" && i.BaseUnit() != f.BaseUnit {
			continue
		}
		return i, nil
	}
	return recipe_dom.Ingredient{}, errors.New(recipe_dom.ErrNotFound)
}

func (imi *InMemoryIngredients) DeleteById(ctx context.Context, id uuid.UUID) error {
	for i, ing := range imi.savedIngredients {
		if id == ing.Id() {
			imi.savedIngredients = append(imi.savedIngredients[0:i], imi.savedIngredients[i+1:]...)
			imi.deletedIngredient = ing
		}
		return nil
	}
	return errors.New(recipe_dom.ErrNotFound)
}

// ---------- Added for tests ------

func (imi *InMemoryIngredients) GetSaved() []recipe_dom.Ingredient {
	return imi.savedIngredients
}

func (imi *InMemoryIngredients) SetIngredientInMem(ingredient recipe_dom.Ingredient) {
	imi.savedIngredients = append(imi.savedIngredients, ingredient)
}

func (imi *InMemoryIngredients) GetLastDeleted() recipe_dom.Ingredient {
	return imi.deletedIngredient

}
