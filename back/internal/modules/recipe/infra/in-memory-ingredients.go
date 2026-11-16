package recipe_infra

import (
	"context"
	"fmt"
	recipe_dom "food/internal/modules/recipe/domain"
	"sync"

	"github.com/google/uuid"
)

type InMemoryIngredients struct {
	m                   *sync.Mutex
	savedIngredients    []*recipe_dom.Ingredient
	expectedIngredient  *recipe_dom.Ingredient
	expectedGetError    bool
	expectedUpdateError bool
	deletedIngredient   *recipe_dom.Ingredient
}

func NewInMemoryIngredients() *InMemoryIngredients {
	return &InMemoryIngredients{
		m:                &sync.Mutex{},
		savedIngredients: make([]*recipe_dom.Ingredient, 0, 10),
	}
}

func (imi *InMemoryIngredients) Save(ctx context.Context, ingredient *recipe_dom.Ingredient) error {
	imi.m.Lock()
	imi.savedIngredients = append(imi.savedIngredients, ingredient.Clone())
	imi.m.Unlock()
	return nil
}

func (imi *InMemoryIngredients) GetById(ctx context.Context, workspaceId uuid.UUID, ingredientId uuid.UUID) (*recipe_dom.Ingredient, error) {
	if imi.expectedGetError {
		return &recipe_dom.Ingredient{}, fmt.Errorf("get ingredient %s : %w", ingredientId.String(), recipe_dom.ErrIngredientNotFound)
	}
	for _, i := range imi.savedIngredients {
		if i.Id() == ingredientId && i.WorkspaceId() == workspaceId {
			return i.Clone(), nil
		}
	}
	return &recipe_dom.Ingredient{}, fmt.Errorf("get ingredient %s : %w", ingredientId.String(), recipe_dom.ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) ExistsByName(ctx context.Context, workspaceId uuid.UUID, name string) (bool, error) {
	for _, i := range imi.savedIngredients {
		if i.WorkspaceId() == workspaceId && i.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

// TODO : check it
func (imi *InMemoryIngredients) Search(ctx context.Context, workspaceId uuid.UUID, f recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	for _, ing := range imi.savedIngredients {
		if ing.WorkspaceId() == workspaceId {
			if f.Name != "" && ing.Name() != f.Name {
				continue
			}
			if f.BaseUnit != "" && ing.BaseUnit() != f.BaseUnit {
				continue
			}
			return ing, nil
		}
	}
	return &recipe_dom.Ingredient{}, recipe_dom.ErrIngredientNotFound
}

func (imi *InMemoryIngredients) Update(ctx context.Context, workspaceId uuid.UUID, ingredient *recipe_dom.Ingredient) error {
	if imi.expectedUpdateError {
		return fmt.Errorf("update ingredient %s : %w", ingredient.Id().String(), recipe_dom.ErrIngredientCouldNotBeUpdated)
	}
	for i, ing := range imi.savedIngredients {
		if ing.Id() == ingredient.Id() {
			imi.m.Lock()
			imi.savedIngredients[i] = ingredient
			imi.m.Unlock()
			return nil
		}
	}
	return fmt.Errorf("update ingredient %s : %w", ingredient.Id().String(), recipe_dom.ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) DeleteById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) error {

	for i, ing := range imi.savedIngredients {
		if workspaceId == ing.WorkspaceId() && id == ing.Id() {
			imi.m.Lock()
			imi.savedIngredients = append(imi.savedIngredients[0:i], imi.savedIngredients[i+1:]...)
			imi.deletedIngredient = ing
			return nil
		}
	}
	return fmt.Errorf("delete ingredient %s : %w", id.String(), recipe_dom.ErrIngredientNotFound)
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
