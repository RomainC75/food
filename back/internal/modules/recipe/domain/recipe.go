package recipe_dom

import (
	"encoding/json"
	"errors"
	"food/utils/displayer"
	"slices"

	"github.com/google/uuid"
)

var (
	ErrInvalidIngredient = "invalid ingredient"
)

type RecipeIngredient struct {
	IngredientID uuid.UUID `json:"ingredient_id"`
	Quantity     Quantity  `json:"qty"`
}

type Recipe struct {
	id          uuid.UUID
	name        string
	ingredients []RecipeIngredient
	workspaceId uuid.UUID
}

func NewRecipe(workspaceId uuid.UUID, id uuid.UUID, name string) *Recipe {
	return &Recipe{
		id:          id,
		name:        name,
		ingredients: make([]RecipeIngredient, 0, 10),
		workspaceId: workspaceId,
	}
}

func (r *Recipe) Id() uuid.UUID {
	return r.id
}

func (r *Recipe) WorkspaceId() uuid.UUID {
	return r.workspaceId
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
	displayer.PrettyDisplay("-after adding : ", r)
	return nil
}

func (r *Recipe) ReplaceIngredients(ris []RecipeIngredient) {
	r.ingredients = ris
}

func (r *Recipe) DeleteIngredient(id uuid.UUID) {
	for index, i := range r.ingredients {
		if i.IngredientID == id {
			r.ingredients = append(r.ingredients[:index], r.ingredients[index+1:]...)
			return
		}
	}
}

func (r *Recipe) DeleteAllIngredients() {
	r.ingredients = make([]RecipeIngredient, 0, 10)
}

func (r *Recipe) IsSameWorkspace(workspaceId uuid.UUID) bool {
	return r.workspaceId == workspaceId
}

// -- EXPORT --

func (r *Recipe) Clone() *Recipe {
	if r == nil {
		return nil
	}
	ingredients := slices.Clone(r.ingredients)
	return &Recipe{
		id:          r.id,
		name:        r.name,
		ingredients: ingredients,
		workspaceId: r.workspaceId,
	}
}

func (r *Recipe) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          uuid.UUID          `json:"id"`
		Name        string             `json:"name"`
		Ingredients []RecipeIngredient `json:"ingredients"`
		WorkspaceId uuid.UUID          `json:"workspace_id"`
	}{
		ID:          r.id,
		Name:        r.name,
		Ingredients: r.ingredients,
		WorkspaceId: r.workspaceId,
	})
}
