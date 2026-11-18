package idaccess_dom

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCouldNotCreateWorkspace = errors.New("could not create workspace")
	ErrWorkspaceNotFound       = errors.New("workspace not found")
)

type Workspaces interface {
	CreateWithMembership(ctx context.Context, workspace *Workspace, membership *Membership) error
	FindById(ctx context.Context, userId uuid.UUID, workspaceId uuid.UUID) (*Workspace, error)
	FindByName(ctx context.Context, userId uuid.UUID, name string) (*Workspace, error)
	GetWorkspaces(ctx context.Context, userId uuid.UUID) ([]*Workspace, error)
	Save(ctx context.Context, userId uuid.UUID, workspace *Workspace) error
	// DeleteById(id uuid.UUID) error
}
