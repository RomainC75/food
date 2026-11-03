package recipe_dom

import (
	"context"
)

type Receipes interface {
	Save(ctx context.Context, r *Receipe) error
}
