package recipe_dom

import (
	"context"

	"github.com/google/uuid"
)

type Recipes interface {
	GetById(ctx context.Context, id uuid.UUID) (*Recipe, error)
	CreateRecipe(ctx context.Context, r *Recipe) error
	UpdateRecipe(ctx context.Context, r *Recipe) error
}
