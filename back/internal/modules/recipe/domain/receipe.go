package recipe_dom

import "github.com/google/uuid"

type Recipe struct {
	id          uuid.UUID
	name        string
	ingredients []*Ingredient
}

func NewRecipe(id uuid.UUID, name string) *Recipe {
	return &Recipe{
		id:          id,
		name:        name,
		ingredients: make([]*Ingredient, 0, 10),
	}
}

func (r *Recipe) AddIngredient(ingredient *Ingredient) {
	r.ingredients = append(r.ingredients, ingredient)
}
