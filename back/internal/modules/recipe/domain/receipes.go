package recipe_dom

import (
	"context"
)

type Recipes interface {
	Save(ctx context.Context, r *Recipe) error
}
