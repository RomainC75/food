package recipe_infra

import (
	"context"
	recipe_dom "food/internal/modules/recipe/domain"
)

type InMemRecipes struct {
	savedRecipes []*recipe_dom.Recipe
}

func NewInMemRecipes() *InMemRecipes {
	return &InMemRecipes{
		savedRecipes: make([]*recipe_dom.Recipe, 0, 10),
	}
}

func (imr *InMemRecipes) Save(ctx context.Context, r *recipe_dom.Recipe) error {
	imr.savedRecipes = append(imr.savedRecipes, r)
	return nil
}

func (imr *InMemRecipes) SavedRecipes(ctx context.Context, r *recipe_dom.Recipe) []*recipe_dom.Recipe {
	return imr.savedRecipes
}
