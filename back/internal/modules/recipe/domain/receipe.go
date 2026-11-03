package recipe_dom

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidIngredient = "invalid ingredient"
)

type RecipeIngredient struct {
	ingredientID uuid.UUID
	quantity     Quantity
}

type Recipe struct {
	id          uuid.UUID
	name        string
	ingredients []RecipeIngredient
}

func NewRecipe(id uuid.UUID, name string) *Recipe {
	return &Recipe{
		id:          id,
		name:        name,
		ingredients: make([]RecipeIngredient, 0, 10),
	}
}

func (r *Recipe) Id() uuid.UUID {
	return r.id
}

func (r *Recipe) AddIngredient(i *Ingredient, q Quantity) error {
	if !i.IsQtyValid(q) {
		return errors.New(ErrInvalidIngredient)
	}

	newRI := RecipeIngredient{
		ingredientID: i.Id(),
		quantity:     q,
	}
	r.ingredients = append(r.ingredients, newRI)
	return nil
}
