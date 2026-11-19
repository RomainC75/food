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
	GetById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) (*Recipe, error)
	GetByName(ctx context.Context, workspaceId uuid.UUID, name string) (*Recipe, error)
	Create(ctx context.Context, r *Recipe) error
	Update(ctx context.Context, r *Recipe) error
	Delete(ctx context.Context, workdspaceId uuid.UUID, id uuid.UUID) error
}
