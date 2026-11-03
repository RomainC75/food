package recipe_dom

import (
	"context"

	"github.com/google/uuid"
)

var (
	ErrNotFound = "ingredient not found"
)

type IngredientFilter struct {
	Name     string
	BaseUnit BaseUnit
}

type Ingredients interface {
	Save(ctx context.Context, ingredient *Ingredient) error
	GetById(ctx context.Context, id uuid.UUID) (*Ingredient, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	Search(ctx context.Context, f IngredientFilter) (*Ingredient, error)
}
