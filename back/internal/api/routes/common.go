package routes

import (
	"food/config"
	idaccess_app "food/internal/modules/id-access/application"
	idaccess_infra "food/internal/modules/id-access/infra"
	recipe_app "food/internal/modules/recipe/application"
	recipe_infra "food/internal/modules/recipe/infra"
	shared_infra "food/internal/shared/infra"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var ()

func ConnectRoutes() *chi.Mux {
	jwtSecret := config.Get().Jwt.Secret
	expiresIn := time.Duration(config.Get().Jwt.ExpiresInMin * int(time.Minute))
	timeDate := shared_infra.NewDateTime()

	ingredientSrv := recipe_app.NewIngredientSrv(
		shared_infra.NewUuidGenerator(),
		recipe_infra.NewInMemoryIngredients().SetIngredientsInMem(),
	)

	userSrv := idaccess_app.NewUserSrv(
		shared_infra.NewUuidGenerator(),
		idaccess_infra.NewBcrypt(),
		idaccess_infra.NewJWT(timeDate, jwtSecret, expiresIn),
		idaccess_infra.NewInMemoryUsers(),
	)

	workspaceSrv := idaccess_app.NewWorkspaceSrv(
		shared_infra.NewUuidGenerator(),
		idaccess_infra.NewInMemWorkspaces(),
		timeDate,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ConnectIngredientRoutes(r, ingredientSrv, userSrv, workspaceSrv)
	ConnectIdAccessRoutes(r, userSrv, workspaceSrv)
	return r
}
