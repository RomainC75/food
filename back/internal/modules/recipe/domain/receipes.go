package recipe_dom

import (
	"context"

	"github.com/google/uuid"
)

var (
	ErrCouldNotUpdateRecipe = "could not update the recipe"
)

type Recipes interface {
	GetById(ctx context.Context, id uuid.UUID) (*Recipe, error)
	Create(ctx context.Context, r *Recipe) error
	Update(ctx context.Context, r *Recipe) error
	Delete(ctx context.Context, id uuid.UUID) error
}
