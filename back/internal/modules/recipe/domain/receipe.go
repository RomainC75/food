package recipe_dom

import "github.com/google/uuid"

type Receipe struct {
	id          uuid.UUID
	name        string
	ingredients []Ingredient
}

func NewReceipe(id uuid.UUID, name string) *Receipe {
	return &Receipe{
		id:          id,
		name:        name,
		ingredients: make([]Ingredient, 0, 10),
	}
}

func (r *Receipe) AddIngredient(ingredient Ingredient) {
	r.ingredients = append(r.ingredients, ingredient)
}
