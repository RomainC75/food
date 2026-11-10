package middlewares

import (
	"context"
	ctrl "food/internal/api/controllers"
	idaccess_dom "food/internal/modules/id-access/domain"
	recipe_dom "food/internal/modules/recipe/domain"
	"net/http"

	"github.com/google/uuid"
)

func WorkspaceMiddleware(authorizer recipe_dom.WorkspaceAuthorizer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId, ok := r.Context().Value(idaccess_dom.TokenUserIdKey).(uuid.UUID)
			if !ok {
				http.Error(w, "id malformed in token", http.StatusBadRequest)
				return
			}
			workspaceIdStr := r.PathValue(ctrl.WorkspaceIdPathValue)
			wsId, err := uuid.Parse(workspaceIdStr)
			if err != nil {
				http.Error(w, "workspace  missing", http.StatusBadRequest)
				return
			}

			err = authorizer.CanCreateInWorkspace(r.Context(), userId, wsId)
			if err != nil {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), idaccess_dom.WorkspaceIdKey, wsId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
