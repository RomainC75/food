package routes

import (
	"context"
	"food/config"
	middw "food/internal/api/middlewares"
	idaccess_app "food/internal/modules/id-access/application"
	idaccess_dom "food/internal/modules/id-access/domain"
	idaccess_infra "food/internal/modules/id-access/infra"
	idaccess_infra_persist "food/internal/modules/id-access/infra/persistence"
	recipe_app "food/internal/modules/recipe/application"
	recipe_infra "food/internal/modules/recipe/infra"
	shared_infra "food/internal/shared/infra"
	utils_must "food/utils/must"
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
	bcrypt := idaccess_infra.NewBcrypt()
	users := idaccess_infra.NewInMemoryUsers()
	users.Create(context.Background(), idaccess_dom.NewUser(uuidGen.Generate(), "fake@email.com", utils_must.Must(bcrypt.Generate("qsdfqsdf"))))
	workspaces := idaccess_infra.NewInMemWorkspaces()

	inMemIngredients := recipe_infra.NewInMemoryIngredients()
	inMemRepises := recipe_infra.NewInMemRecipes()

	// ======= IdAccess services ======
	userSrv := idaccess_app.NewUserSrv(
		uuidGen,
		bcrypt,
		idaccess_infra.NewJWT(
			timeDate,
			jwtRefreshSecret,
			jwtAccessSecret,
			refreshExpiresIn,
			accessExpiresIn,
		),
		users,
		workspaces,
	)
	idaccess_infra_persist.PopulateUsers(userSrv)

	workspaceSrv := idaccess_app.NewWorkspaceSrv(
		uuidGen,
		workspaces,
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
