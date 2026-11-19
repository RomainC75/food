package recipe_infra

import (
	"context"
	"fmt"
	recipe_dom "food/internal/modules/recipe/domain"
	"sync"

	"github.com/google/uuid"
)

var ()

type RecipeRepoErrorsToSet struct {
	UpdateError bool
}

type InMemRecipes struct {
	m              sync.RWMutex
	savedRecipes   map[uuid.UUID]*recipe_dom.Recipe
	expectedErrors RecipeRepoErrorsToSet
}

func NewInMemRecipes() *InMemRecipes {
	return &InMemRecipes{
		m:            sync.RWMutex{},
		savedRecipes: make(map[uuid.UUID]*recipe_dom.Recipe, 10),
	}
}

func (imr *InMemRecipes) Create(ctx context.Context, r *recipe_dom.Recipe) error {
	imr.m.Lock()
	defer imr.m.Unlock()

	imr.savedRecipes[r.Id()] = r.Clone()
	return nil
}

func (imr *InMemRecipes) GetById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) (*recipe_dom.Recipe, error) {
	imr.m.Lock()
	defer imr.m.Unlock()

	foundRecipe, ok := imr.getIngredientInWorkspace(workspaceId, id)

	if !ok {
		return nil, fmt.Errorf("get recipe by id %s : %w", id.String(), recipe_dom.ErrRecipeNotFound)
	}

	return foundRecipe.Clone(), nil
}

func (imr *InMemRecipes) GetByName(ctx context.Context, workspaceId uuid.UUID, name string) (*recipe_dom.Recipe, error) {
	imr.m.Lock()
	defer imr.m.Unlock()

	for _, r := range imr.savedRecipes {
		if r.Name() == name {
			return r.Clone(), nil
		}
	}
	return nil, fmt.Errorf("could not get recipe by name %s : %w", name, recipe_dom.ErrRecipeNotFound)
}

func (imr *InMemRecipes) Update(ctx context.Context, r *recipe_dom.Recipe) error {
	if imr.expectedErrors.UpdateError {
		return fmt.Errorf("update ingredient %s : %w", r.Id().String(), recipe_dom.ErrCouldNotUpdateRecipe)
	}

	imr.m.Lock()
	defer imr.m.Unlock()
	_, ok := imr.getIngredientInWorkspace(r.WorkspaceId(), r.Id())
	if !ok {
		return fmt.Errorf("update ingredient %s : %w", r.Id().String(), recipe_dom.ErrRecipeNotFound)
	}

	imr.savedRecipes[r.Id()] = r.Clone()

	return nil
	// !! for the real one
	// 	Updates the recipe.
	// Deletes its existing ingredient relationships.
	// Inserts the current ingredient collection.
}

func (imr *InMemRecipes) Delete(ctx context.Context, workdspaceId uuid.UUID, id uuid.UUID) error {
	imr.m.Lock()
	defer imr.m.Unlock()
	_, ok := imr.getIngredientInWorkspace(workdspaceId, id)
	if !ok {
		return fmt.Errorf("update ingredient %s : %w", id.String(), recipe_dom.ErrRecipeNotFound)
	}
	delete(imr.savedRecipes, id)

	return nil
}

func (imr *InMemRecipes) getIngredientInWorkspace(workspaceId uuid.UUID, id uuid.UUID) (*recipe_dom.Recipe, bool) {
	found, ok := imr.savedRecipes[id]
	if !ok {
		return nil, false
	}

	isSameWorkspace := found.IsSameWorkspace(workspaceId)
	if !isSameWorkspace {
		return nil, false
	}
	return found, true
}

// == TEST ==

func (imr *InMemRecipes) GetSavedRecipes() []*recipe_dom.Recipe {
	r := make([]*recipe_dom.Recipe, 0, len(imr.savedRecipes))
	for _, v := range imr.savedRecipes {
		r = append(r, v)
	}
	return r
}

func (imr *InMemRecipes) SetExpectedErrors(ee RecipeRepoErrorsToSet) {
	imr.expectedErrors = ee
}
