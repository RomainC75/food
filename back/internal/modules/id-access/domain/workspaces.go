package idaccess_dom

import (
	"context"
)

type Workspaces interface {
	CreateWithMembership(ctx context.Context, workspace *Workspace, membership *Membership) error
	// Save() error
	// DeleteById(id uuid.UUID) error
}
