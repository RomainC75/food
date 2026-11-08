package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ConnectRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ConnectIngredientRoutes(r)
	ConnectIdAccessRoutes(r)
	return r
}
