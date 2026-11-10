package platform_http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

var (
	userIdKey    = "user_id"
	emailKey     = "email"
	workspaceKey = "workspace_id"
)

// ==== UserId ===

func WithUserId(ctx context.Context, userId uuid.UUID) context.Context {
	return context.WithValue(ctx, userIdKey, userId)
}

func GetUserIdFromContext(ctx context.Context) (uuid.UUID, bool) {
	userIdStr := ctx.Value(userIdKey)
	userId, ok := userIdStr.(uuid.UUID)
	return userId, ok

}

// ==== Email ===

func WithUserEmail(ctx context.Context, userEmail string) context.Context {
	return context.WithValue(ctx, emailKey, userEmail)
}

func GetEmailFromContext(r *http.Request) string {
	return r.Context().Value(emailKey).(string)
}

// ==== Workspaceid ===

func WithWorkspaceId(ctx context.Context, workspaceId uuid.UUID) context.Context {
	return context.WithValue(ctx, workspaceKey, workspaceId)
}

func GetWorkspaceIdFromContext(ctx context.Context) (uuid.UUID, bool) {
	workspaceIdStr := ctx.Value(workspaceKey)
	workspaceId, ok := workspaceIdStr.(uuid.UUID)
	return workspaceId, ok

}
