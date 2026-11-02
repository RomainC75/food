package recipe_infra

import recipe_dom "food/internal/modules/recipe/domain"

type InMemoryIngredients struct {
	savedIngredients []recipe_dom.Ingredient
}

func NewInMemoryIngredients() *InMemoryIngredients {
	return &InMemoryIngredients{}
}

func (imi *InMemoryIngredients) Save(ingredient recipe_dom.Ingredient) error {
	imi.savedIngredients = append(imi.savedIngredients, ingredient)
	return nil
}

func (imi *InMemoryIngredients) GetSaved() []recipe_dom.Ingredient {
	return imi.savedIngredients
}
