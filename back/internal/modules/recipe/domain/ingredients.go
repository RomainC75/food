package recipe_dom

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrIngredientNotFound          = errors.New("ingredient not found")
	ErrIngredientCouldNotBeUpdated = errors.New("ingredient could not be updated")
)

type IngredientFilter struct {
	BaseUnit BaseUnit
	Name     string
}

type Ingredients interface {
	Save(ctx context.Context, ingredient *Ingredient) error
	GetById(ctx context.Context, workspaceId uuid.UUID, ingredientId uuid.UUID) (*Ingredient, error)
	ExistsByName(ctx context.Context, workspaceId uuid.UUID, name string) (bool, error)
	Search(ctx context.Context, workspaceId uuid.UUID, f IngredientFilter) (*Ingredient, error)
	Update(ctx context.Context, workspaceId uuid.UUID, ingredient *Ingredient) error
	DeleteById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) error
}
