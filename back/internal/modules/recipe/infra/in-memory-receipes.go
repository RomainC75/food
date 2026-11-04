package recipe_infra

import (
	"context"
	"errors"
	"fmt"
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
	fmt.Println("-- get by id : ", imr.savedRecipes, id)

	for _, r := range imr.savedRecipes {
		if r.Id() == id {
			return r.Clone(), nil
		}
	}
	return nil, errors.New(ErrRecipeNotFound)
}

func (imr *InMemRecipes) Update(ctx context.Context, r *recipe_dom.Recipe) error {
	// ! pointer -> test probles
	for i, sr := range imr.savedRecipes {
		if sr.Id() == r.Id() {
			imr.savedRecipes[i] = r.Clone()
			return nil
		}
	}
	return errors.New(ErrRecipeNotFound)
	// !! for the real one
	// 	Updates the recipe.
	// Deletes its existing ingredient relationships.
	// Inserts the current ingredient collection.
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
