package recipe_dom

type Ingredients interface {
	Save(ingredient Ingredient) error
}
