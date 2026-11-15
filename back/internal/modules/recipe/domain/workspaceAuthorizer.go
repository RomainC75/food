package recipe_dom

import (
	"context"

	"github.com/google/uuid"
)

type WorkspaceAuthorizer interface {
	IsMember(
		ctx context.Context,
		userID uuid.UUID,
		workspaceID uuid.UUID,
	) error
}
