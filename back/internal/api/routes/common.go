package routes

import (
	"food/config"
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
	jwtSecret := config.Get().Jwt.Secret
	expiresIn := time.Duration(config.Get().Jwt.ExpiresInMin * int(time.Minute))
	timeDate := shared_infra.NewDateTime()
	uuidGen := shared_infra.NewInMemUuidGen()

	ingredientSrv := recipe_app.NewIngredientSrv(
		uuidGen,
		recipe_infra.NewInMemoryIngredients().SetIngredientsInMem(),
	)

	userSrv := idaccess_app.NewUserSrv(
		uuidGen,
		idaccess_infra.NewBcrypt(),
		idaccess_infra.NewJWT(timeDate, jwtSecret, expiresIn),
		idaccess_infra.NewInMemoryUsers(),
	)
	idaccess_infra_persist.PopulateUsers(uuidGen, userSrv)

	workspaceSrv := idaccess_app.NewWorkspaceSrv(
		uuidGen,
		idaccess_infra.NewInMemWorkspaces(),
		timeDate,
	)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	ConnectIngredientRoutes(r, ingredientSrv, userSrv, workspaceSrv)
	ConnectIdAccessRoutes(r, userSrv, workspaceSrv)
	return r
}
