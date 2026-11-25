package routes

import (
	ctrl "food/internal/api/controllers"
	idaccess_app "food/internal/modules/id-access/application"

	"github.com/go-chi/chi/v5"
)

func ConnectIdAccessRoutes(r chi.Router, userSrv *idaccess_app.UserSrv, workspaceSrv *idaccess_app.WorkspaceSrv) {
	idAccessCtrl := ctrl.NewIdAccessCtrl(userSrv)
	workspaceCtrl := ctrl.NewWorkspaceCtrl(workspaceSrv)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", idAccessCtrl.Signup)
		r.Post("/login", idAccessCtrl.Login)
		r.Get("/refresh", idAccessCtrl.Refresh)

		r.Group(func(r chi.Router) {
			r.Use(idAccessCtrl.Verify)
			r.Get("/verify", idAccessCtrl.ReturnAuthData)
			r.Post("/workspace", workspaceCtrl.Create)
		})
	})
}
