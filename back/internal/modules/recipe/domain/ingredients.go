package recipe_dom

import (
	"context"

	"github.com/google/uuid"
)

var (
	ErrIngredientNotFound          = "ingredient not found"
	ErrIngredientCouldNotBeUpdated = "ingredient could not be updated"
)

type IngredientFilter struct {
	Id       *uuid.UUID
	Name     string
	BaseUnit BaseUnit
}

type Ingredients interface {
	Save(ctx context.Context, ingredient *Ingredient) error
	GetById(ctx context.Context, workspaceId uuid.UUID, ingredientId uuid.UUID) (*Ingredient, error)
	ExistsByName(ctx context.Context, workspaceId uuid.UUID, name string) (bool, error)
	Search(ctx context.Context, f IngredientFilter) (*Ingredient, error)
	Update(ctx context.Context, workspaceId uuid.UUID, ingredient *Ingredient) error
	DeleteById(ctx context.Context, workspaceId uuid.UUID, id uuid.UUID) error
}
