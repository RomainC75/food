package recipe_infra

import (
	"context"
	"fmt"
	recipe_dom "food/internal/modules/recipe/domain"
	"sync"

	"github.com/google/uuid"
)

type InMemoryIngredients struct {
	m                   sync.RWMutex
	savedIngredients    map[uuid.UUID]*recipe_dom.Ingredient
	expectedIngredient  *recipe_dom.Ingredient
	expectedGetError    bool
	expectedUpdateError bool
	deletedIngredient   *recipe_dom.Ingredient
}

func NewInMemoryIngredients() *InMemoryIngredients {
	return &InMemoryIngredients{
		m:                sync.RWMutex{},
		savedIngredients: make(map[uuid.UUID]*recipe_dom.Ingredient, 10),
	}
}

func (imi *InMemoryIngredients) Save(ctx context.Context, ingredient *recipe_dom.Ingredient) error {
	imi.m.Lock()
	defer imi.m.Unlock()

	imi.savedIngredients[ingredient.Id()] = ingredient.Clone()
	return nil
}

func (imi *InMemoryIngredients) GetById(ctx context.Context, workspaceId uuid.UUID, ingredientId uuid.UUID) (*recipe_dom.Ingredient, error) {
	if imi.expectedGetError {
		return &recipe_dom.Ingredient{}, fmt.Errorf("get ingredient %s : %w", ingredientId.String(), recipe_dom.ErrIngredientNotFound)
	}
	imi.m.RLock()
	defer imi.m.RUnlock()

	for _, i := range imi.savedIngredients {
		if i.Id() == ingredientId && i.WorkspaceId() == workspaceId {
			return i.Clone(), nil
		}
	}
	return &recipe_dom.Ingredient{}, fmt.Errorf("get ingredient %s : %w", ingredientId.String(), recipe_dom.ErrIngredientNotFound)
}

func (imi *InMemoryIngredients) ExistsByName(ctx context.Context, workspaceId uuid.UUID, name string) (bool, error) {
	imi.m.RLock()
	defer imi.m.RUnlock()

	for _, i := range imi.savedIngredients {
		if i.IsSameWorkspace(workspaceId) && i.Name() == name {
			return true, nil
		}
	}
	return false, nil
}

// TODO : check it
func (imi *InMemoryIngredients) Search(ctx context.Context, workspaceId uuid.UUID, f recipe_dom.IngredientFilter) (*recipe_dom.Ingredient, error) {
	imi.m.RLock()
	defer imi.m.RUnlock()

	for _, ing := range imi.savedIngredients {
		if ing.WorkspaceId() == workspaceId {
			if f.Name != "" && ing.Name() != f.Name {
				continue
			}
			if f.BaseUnit != "" && ing.BaseUnit() != f.BaseUnit {
				continue
			}
			return ing.Clone(), nil
		}
	}
	return &recipe_dom.Ingredient{}, recipe_dom.ErrIngredientNotFound
}

func (imi *InMemoryIngredients) Update(ctx context.Context, workspaceId uuid.UUID, ingredient *recipe_dom.Ingredient) error {
	if imi.expectedUpdateError {
		return fmt.Errorf("update ingredient %s : %w", ingredient.Id().String(), recipe_dom.ErrIngredientCouldNotBeUpdated)
	}

	imi.m.Lock()
	defer imi.m.Unlock()

	_, ok := imi.getIngredientInWorkspace(workspaceId, ingredient.Id())
	if !ok {
		return fmt.Errorf("update ingredient %s : %w", ingredient.Id().String(), recipe_dom.ErrIngredientNotFound)
	}

	imi.savedIngredients[ingredient.Id()] = ingredient.Clone()

	return nil
}

func (imi *InMemoryIngredients) DeleteById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) error {
	imi.m.Lock()
	defer imi.m.Unlock()

	found, ok := imi.getIngredientInWorkspace(workspaceId, id)
	if !ok {
		return fmt.Errorf("update ingredient %s : %w", id.String(), recipe_dom.ErrIngredientNotFound)
	}

	imi.deletedIngredient = found
	delete(imi.savedIngredients, id)

	return nil
}

func (imi *InMemoryIngredients) getIngredientInWorkspace(workspaceId uuid.UUID, id uuid.UUID) (*recipe_dom.Ingredient, bool) {
	found, ok := imi.savedIngredients[id]
	if !ok {
		return nil, false
	}

	isSameWorkspace := found.IsSameWorkspace(workspaceId)
	if !isSameWorkspace {
		return nil, false
	}
	return found, true
}

// ---------- Added for tests ------

func (imi *InMemoryIngredients) GetSaved() map[uuid.UUID]*recipe_dom.Ingredient {
	return imi.savedIngredients
}

func (imi *InMemoryIngredients) SetIngredientsInMem(ings ...*recipe_dom.Ingredient) *InMemoryIngredients {
	for _, ing := range ings {
		imi.savedIngredients[ing.Id()] = ing.Clone()
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
