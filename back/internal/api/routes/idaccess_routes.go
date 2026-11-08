package routes

import (
	ctrl "food/internal/api/controllers"

	"github.com/go-chi/chi/v5"
)

func ConnectIdAccessRoutes(r *chi.Mux) {
	idAccessCtrl := ctrl.NewIdAccessCtrl()

	r.Route("/id-access", func(r chi.Router) {
		r.Post("/signup", idAccessCtrl.Signup)

		r.Post("/signin", idAccessCtrl.Signin)

	})
	r.Route("/verify", func(r chi.Router) {
		r.Use(idAccessCtrl.Verify)
		r.Get("/", idAccessCtrl.ReturnAuthData)

	})
}
