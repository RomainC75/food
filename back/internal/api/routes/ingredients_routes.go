package routes

import (
	ctrl "food/internal/api/controllers"
	"food/internal/api/middlewares"
	idaccess_app "food/internal/modules/id-access/application"
	recipe_app "food/internal/modules/recipe/application"
	recipe_dom "food/internal/modules/recipe/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ConnectIngredientRoutes(r *chi.Mux, ingredientSrv *recipe_app.IngredientSrv, userSrv *idaccess_app.UserSrv, authorizer recipe_dom.WorkspaceAuthorizer) {
	ingredientCtrl := ctrl.NewIngredientCtrl(ingredientSrv)
	userCtrl := ctrl.NewIdAccessCtrl(userSrv)
	r.Route("/workspace/{workspaceId}/ingredients", func(r chi.Router) {
		r.Use(userCtrl.Verify)
		r.Use(middlewares.WorkspaceMiddleware(authorizer))
		r.Post("/", ingredientCtrl.CreateIncredient)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ingredientCtrl.Search(w, r)
			// w.Write([]byte("welcome"))
		})

	})
}
