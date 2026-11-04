package recipe_infra

import (
	"context"
	"errors"
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/google/uuid"
)

var (
	ErrRecipeNotFound = "recipe not found"
)

type InMemRecipes struct {
	savedRecipes []*recipe_dom.Recipe
}

func NewInMemRecipes() *InMemRecipes {
	return &InMemRecipes{
		savedRecipes: make([]*recipe_dom.Recipe, 0, 10),
	}
}

func (imr *InMemRecipes) Create(ctx context.Context, r *recipe_dom.Recipe) error {
	imr.savedRecipes = append(imr.savedRecipes, r)
	return nil
}

func (imr *InMemRecipes) GetById(ctx context.Context, id uuid.UUID) (*recipe_dom.Recipe, error) {
	for _, r := range imr.savedRecipes {
		if r.Id() == id {
			return r, nil
		}
	}
	return nil, errors.New(ErrRecipeNotFound)
}

func (imr *InMemRecipes) Update(ctx context.Context, r *recipe_dom.Recipe) error {
	// ** auto - save : pointer

	// !! for the real one
	// 	Updates the recipe.
	// Deletes its existing ingredient relationships.
	// Inserts the current ingredient collection.
	return nil
}

func (imr *InMemRecipes) Delete(ctx context.Context, id uuid.UUID) error {
	for i, r := range imr.savedRecipes {
		if r.Id() == id {
			imr.savedRecipes = append(imr.savedRecipes[:i], imr.savedRecipes[i+1:]...)
			return nil
		}
	}
	return errors.New(ErrRecipeNotFound)
}

// == TEST ==

func (imr *InMemRecipes) GetSavedRecipes() []*recipe_dom.Recipe {
	return imr.savedRecipes
}

func (imr *InMemRecipes) DeleteById(id uuid.UUID) error {
	for i, r := range imr.savedRecipes {
		if r.Id() == id {
			imr.savedRecipes = append(imr.savedRecipes[:i], imr.savedRecipes[i+1:]...)
			return nil
		}
	}
	return errors.New(ErrRecipeNotFound)
}
