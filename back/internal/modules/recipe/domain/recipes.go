package recipe_dom

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCouldNotUpdateRecipe = errors.New("could not update the recipe")
	ErrRecipeNotFound       = errors.New("recipe not found")
)

type Recipes interface {
	GetById(ctx context.Context, id uuid.UUID) (*Recipe, error)
	Create(ctx context.Context, r *Recipe) error
	Update(ctx context.Context, r *Recipe) error
	Delete(ctx context.Context, id uuid.UUID) error
}
