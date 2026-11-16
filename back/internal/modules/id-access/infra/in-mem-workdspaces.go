package idaccess_infra

import (
	"context"
	"errors"
	idaccess_dom "food/internal/modules/id-access/domain"
	"slices"
	"sync"

	"github.com/google/uuid"
)

const (
	ErrWorkspaceNotFound    = "workspace not found"
	ErrNoWorkspaceAvailable = "no workspace available"
)

type InMemWorkspaces struct {
	m               sync.RWMutex
	workspaces      map[uuid.UUID]*idaccess_dom.Workspace
	memberships     map[uuid.UUID]*idaccess_dom.Membership
	isErrorExpected bool
}

func NewInMemWorkspaces() *InMemWorkspaces {
	return &InMemWorkspaces{
		m:           sync.RWMutex{},
		workspaces:  make(map[uuid.UUID]*idaccess_dom.Workspace, 10),
		memberships: make(map[uuid.UUID]*idaccess_dom.Membership, 10),
	}
}

func (imw *InMemWorkspaces) CreateWithMembership(ctx context.Context, workspace *idaccess_dom.Workspace, membership *idaccess_dom.Membership) error {
	if imw.isErrorExpected {
		return idaccess_dom.ErrCouldNotCreateWorkspace
	}
	imw.m.Lock()
	imw.workspaces[workspace.Id()] = workspace.Clone()
	imw.memberships[membership.Id()] = membership.Clone()
	imw.m.Unlock()
	return nil
}

func (imw *InMemWorkspaces) FindById(ctx context.Context, userId uuid.UUID, workspaceId uuid.UUID) (*idaccess_dom.Workspace, error) {
	imw.m.Lock()
	w, ok := imw.workspaces[workspaceId]
	imw.m.Unlock()
	if !ok {
		return nil, errors.New(ErrWorkspaceNotFound)
	}
	imw.m.Lock()
	for _, m := range imw.memberships {
		if m.IsSameWorkspace(w.Id()) && m.IsSameUser(userId) {
			return w.Clone(), nil
		}
	}
	imw.m.Unlock()
	return nil, errors.New(ErrWorkspaceNotFound)
}

func (imw *InMemWorkspaces) GetWorkspaces(ctx context.Context, userId uuid.UUID) ([]*idaccess_dom.Workspace, error) {
	wpIds := make([]uuid.UUID, 0, 10)
	imw.m.RLock()
	for _, m := range imw.memberships {
		if m.IsSameUser(userId) {
			wpIds = append(wpIds, m.WorkspaceId())
		}
	}
	imw.m.RLock()
	if len(wpIds) == 0 {
		return []*idaccess_dom.Workspace{}, nil
	}
	wp := make([]*idaccess_dom.Workspace, 0, 10)
	imw.m.RLock()
	for _, w := range imw.workspaces {
		if slices.Contains(wpIds, w.Id()) {
			wp = append(wp, w)
		}
	}
	imw.m.RUnlock()
	return wp, nil

}

// ===== TEST ===

func (imw *InMemWorkspaces) SavedMemberships() map[uuid.UUID]*idaccess_dom.Membership {
	return imw.memberships
}

func (imw *InMemWorkspaces) SavedWorkspaces() map[uuid.UUID]*idaccess_dom.Workspace {
	return imw.workspaces
}

func (imw *InMemWorkspaces) SetIsErrorExpected(isErrorExpected bool) {
	imw.isErrorExpected = isErrorExpected
}
