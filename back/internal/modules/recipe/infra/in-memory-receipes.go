package recipe_infra

import (
	"context"
	recipe_dom "food/internal/modules/recipe/domain"
)

type InMemReceipes struct {
	savedReceipes []*recipe_dom.Receipe
}

func NewInMemReceipes() *InMemReceipes {
	return &InMemReceipes{
		savedReceipes: make([]*recipe_dom.Receipe, 0, 10),
	}
}

func (imr *InMemReceipes) Save(ctx context.Context, r *recipe_dom.Receipe) error {
	imr.savedReceipes = append(imr.savedReceipes, r)
	return nil
}

func (imr *InMemReceipes) SavedReceipes(ctx context.Context, r *recipe_dom.Receipe) []*recipe_dom.Receipe {
	return imr.savedReceipes
}
