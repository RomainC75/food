package recipe_dom

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidIngredient = "invalid ingredient"
)

type RecipeIngredient struct {
	IngredientID uuid.UUID
	Quantity     Quantity
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

func (r *Recipe) SetName(name string) {
	r.name = name
}

func (r *Recipe) AddIngredient(i *Ingredient, q Quantity) error {
	if !i.IsQtyValid(q) {
		return errors.New(ErrInvalidIngredient)
	}

	newRI := RecipeIngredient{
		IngredientID: i.Id(),
		Quantity:     q,
	}
	r.ingredients = append(r.ingredients, newRI)
	return nil
}

func (r *Recipe) UpdateIngredients(ris []RecipeIngredient) {
	newList := make([]RecipeIngredient, 0, len(r.ingredients))
	for _, i := range r.ingredients {
		for _, nri := range ris {
			if i.IngredientID == nri.IngredientID {
				i.Quantity = nri.Quantity
				newList = append(newList, i)
			}
		}
	}
	r.ingredients = newList
}

func (r *Recipe) DeleteIngredient(id uuid.UUID) {
	for index, i := range r.ingredients {
		if i.IngredientID == id {
			r.ingredients = append(r.ingredients[:index], r.ingredients[index+1:]...)
			return
		}
	}
}
