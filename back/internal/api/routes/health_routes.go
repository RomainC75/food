package routes

import (
	ctrl "food/internal/api/controllers"

	"github.com/go-chi/chi/v5"
)

func ConnectHealthRoutes(r *chi.Mux) {
	healthCtrl := ctrl.NewHealthCtrl()

	r.Route("/health", func(r chi.Router) {
		r.Get("/", healthCtrl.Health)
	})
}
