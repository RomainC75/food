package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ConnectIngredientRoutes(r *chi.Mux) {
	r.Route("/ingredient", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
	})
}
