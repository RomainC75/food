package routes

import (
	ctrl "food/internal/api/controllers"
	"food/internal/api/middlewares"
	idaccess_app "food/internal/modules/id-access/application"
	recipe_app "food/internal/modules/recipe/application"
	recipe_dom "food/internal/modules/recipe/domain"

	"github.com/go-chi/chi/v5"
)

func ConnectRecipeRoutes(r chi.Router, recipeSrv *recipe_app.RecipeSrv, userSrv *idaccess_app.UserSrv, authorizer recipe_dom.WorkspaceAuthorizer) {
	recipeCtrl := ctrl.NewRecipeCtrl(recipeSrv)
	userCtrl := ctrl.NewIdAccessCtrl(userSrv)
	r.Route("/workspace/{workspaceId}/recipes", func(r chi.Router) {
		r.Use(userCtrl.Verify)
		r.Use(middlewares.WorkspaceMiddleware(authorizer))

		r.Post("/", recipeCtrl.Create)
		r.Put("/", recipeCtrl.Update)
		// r.Get("/", recipeCtrl.Search)
		// r.Delete("/{ingredient-id}", recipeCtrl.Delete)

	})
}
