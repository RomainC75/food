package routes

import (
	"food/config"
	middw "food/internal/api/middlewares"
	idaccess_app "food/internal/modules/id-access/application"
	idaccess_infra "food/internal/modules/id-access/infra"
	idaccess_infra_persist "food/internal/modules/id-access/infra/persistence"
	recipe_app "food/internal/modules/recipe/application"
	recipe_infra "food/internal/modules/recipe/infra"
	shared_infra "food/internal/shared/infra"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ConnectRoutes() *chi.Mux {
	jwtRefreshSecret := config.Get().Jwt.RefreshSecret
	jwtAccessSecret := config.Get().Jwt.AccessSecret
	refreshExpiresIn := time.Duration(config.Get().Jwt.RefreshExpiresInDays * int(time.Hour) * 24)
	accessExpiresIn := time.Duration(config.Get().Jwt.AccessExpiresInMin * int(time.Minute))
	timeDate := shared_infra.NewDateTime()
	uuidGen := shared_infra.NewUuidGenerator()

	inMemIngredients := recipe_infra.NewInMemoryIngredients().SetIngredientsInMem()
	inMemRepises := recipe_infra.NewInMemRecipes()

	// ======= IdAccess services ======
	userSrv := idaccess_app.NewUserSrv(
		uuidGen,
		idaccess_infra.NewBcrypt(),
		idaccess_infra.NewJWT(timeDate, jwtRefreshSecret, jwtAccessSecret, refreshExpiresIn, accessExpiresIn),
		idaccess_infra.NewInMemoryUsers(),
	)
	idaccess_infra_persist.PopulateUsers(userSrv)

	workspaceSrv := idaccess_app.NewWorkspaceSrv(
		uuidGen,
		idaccess_infra.NewInMemWorkspaces(),
		timeDate,
	)

	// ======= Recipe services ======
	ingredientSrv := recipe_app.NewIngredientSrv(
		uuidGen,
		inMemIngredients,
	)

	recipeSrv := recipe_app.NewRecipeSrv(
		uuidGen,
		inMemIngredients,
		inMemRepises,
	)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middw.Cors)

	ConnectHealthRoutes(r)
	r.Route("/api", func(r chi.Router) {
		ConnectIdAccessRoutes(r, userSrv, workspaceSrv)
		ConnectIngredientRoutes(r, ingredientSrv, userSrv, workspaceSrv)
		ConnectRecipeRoutes(r, recipeSrv, userSrv, workspaceSrv)
	})

	return r
}
