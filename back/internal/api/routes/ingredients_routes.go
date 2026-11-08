package routes

import (
	ctrl "food/internal/api/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ConnectIngredientRoutes(r *chi.Mux) {
	ingredientCtrl := ctrl.NewIngredientCtrl()

	r.Route("/ingredient", func(r chi.Router) {
		r.Use()
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ingredientCtrl.Search(w, r)
			// w.Write([]byte("welcome"))
		})

	})
}
