package idaccess_dom

import (
	"context"
)

const (
	ErrCouldNotCreateWorkspace = "could not create workspace"
)

type Workspaces interface {
	CreateWithMembership(ctx context.Context, workspace *Workspace, membership *Membership) error
	// Save() error
	// DeleteById(id uuid.UUID) error
}
