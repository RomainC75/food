package routes

import (
	ctrl "food/internal/api/controllers"
	"food/internal/api/middlewares"
	idaccess_app "food/internal/modules/id-access/application"
	recipe_app "food/internal/modules/recipe/application"
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/go-chi/chi/v5"
)

func ConnectIngredientRoutes(r *chi.Mux, ingredientSrv *recipe_app.IngredientSrv, userSrv *idaccess_app.UserSrv, authorizer recipe_dom.WorkspaceAuthorizer) {
	ingredientCtrl := ctrl.NewIngredientCtrl(ingredientSrv)
	userCtrl := ctrl.NewIdAccessCtrl(userSrv)
	r.Route("/workspace/{workspaceId}/ingredients", func(r chi.Router) {
		r.Use(userCtrl.Verify)
		r.Use(middlewares.WorkspaceMiddleware(authorizer))

		r.Post("/", ingredientCtrl.CreateIncredient)
		r.Get("/", ingredientCtrl.Search)
		r.Delete("/{ingredient-id}", ingredientCtrl.Delete)

	})
}
