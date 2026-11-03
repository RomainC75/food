package recipe_dom

import (
	"context"
)

type Recipes interface {
	CreateRecipe(ctx context.Context, r *Recipe) error
	UpdateRecipe(ctx context.Context, r *Recipe) error
}
