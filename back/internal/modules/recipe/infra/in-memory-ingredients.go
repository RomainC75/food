package recipe_infra

import (
	"context"
	"errors"
	recipe_dom "food/internal/modules/recipe/domain"
	"log"

	"github.com/google/uuid"
)

type InMemoryIngredients struct {
	savedIngredients    []*recipe_dom.Ingredient
	expectedIngredient  *recipe_dom.Ingredient
	expectedGetError    bool
	expectedUpdateError bool
	deletedIngredient   *recipe_dom.Ingredient
}

func NewInMemoryIngredients() *InMemoryIngredients {
	return &InMemoryIngredients{
		savedIngredients: make([]*recipe_dom.Ingredient, 0, 10),
	}
}

func (imi *InMemoryIngredients) Save(ctx context.Context, ingredient *recipe_dom.Ingredient) error {
	imi.savedIngredients = append(imi.savedIngredients, ingredient.Clone())
	log.Println(" after save ---------", imi.savedIngredients)
	return nil
}

func (imi *InMemoryIngredients) GetById(ctx context.Context, workspaceId uuid.UUID, ingredientId uuid.UUID) (*recipe_dom.Ingredient, error) {
	if imi.expectedGetError {
		return &recipe_dom.Ingredient{}, errors.New(recipe_dom.ErrIngredientNotFound)
	}
	for _, i := range imi.savedIngredients {
		if i.Id() == ingredientId && i.WorkspaceId() == workspaceId {
			return i.Clone(), nil
		}
	}
	return &recipe_dom.Ingredient{}, errors.New(recipe_dom.ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) ExistsByName(ctx context.Context, workspaceId uuid.UUID, name string) (bool, error) {
	for _, i := range imi.savedIngredients {
		if i.WorkspaceId() == workspaceId && i.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

func (imi *InMemoryIngredients) Search(ctx context.Context, f recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	for _, ing := range imi.savedIngredients {
		if f.Id != nil && ing.Id() != *f.Id {
			continue
		}
		if f.Name != "" && ing.Name() != f.Name {
			continue
		}
		if f.BaseUnit != "" && ing.BaseUnit() != f.BaseUnit {
			continue
		}
		return ing, nil
	}
	return &recipe_dom.Ingredient{}, errors.New(recipe_dom.ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) Update(ctx context.Context, workspaceId uuid.UUID, ingredient *recipe_dom.Ingredient) error {
	if imi.expectedUpdateError {
		return errors.New(recipe_dom.ErrIngredientCouldNotBeUpdated)
	}
	for i, ing := range imi.savedIngredients {
		if ing.Id() == ingredient.Id() {
			imi.savedIngredients[i] = ingredient
			return nil
		}
	}
	return errors.New(recipe_dom.ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) DeleteById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) error {
	for i, ing := range imi.savedIngredients {
		if workspaceId == ing.WorkspaceId() && id == ing.Id() {
			imi.savedIngredients = append(imi.savedIngredients[0:i], imi.savedIngredients[i+1:]...)
			imi.deletedIngredient = ing
		}
		return nil
	}
	return errors.New(recipe_dom.ErrIngredientNotFound)
}

// ---------- Added for tests ------

func (imi *InMemoryIngredients) GetSaved() []*recipe_dom.Ingredient {
	return imi.savedIngredients
}

func (imi *InMemoryIngredients) SetIngredientsInMem(ings ...*recipe_dom.Ingredient) *InMemoryIngredients {
	for _, ing := range ings {
		imi.savedIngredients = append(imi.savedIngredients, ing)
	}
	return imi
}

func (imi *InMemoryIngredients) GetLastDeleted() *recipe_dom.Ingredient {
	return imi.deletedIngredient

}

func (imi *InMemoryIngredients) SetExpectedErrors(expectedGetError bool, expectedUpdateError bool) {
	imi.expectedGetError = expectedGetError
	imi.expectedUpdateError = expectedUpdateError
}
