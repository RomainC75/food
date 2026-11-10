package idaccess_dom

import (
	"context"

	"github.com/google/uuid"
)

const (
	ErrCouldNotCreateWorkspace = "could not create workspace"
)

type Workspaces interface {
	CreateWithMembership(ctx context.Context, workspace *Workspace, membership *Membership) error
	FindById(ctx context.Context, userId uuid.UUID, workspaceId uuid.UUID) (*Workspace, error)
	// Save() error
	// DeleteById(id uuid.UUID) error
}
