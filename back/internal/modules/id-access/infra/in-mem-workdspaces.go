package idaccess_infra

import (
	"context"
	"errors"
	"fmt"
	idaccess_dom "food/internal/modules/id-access/domain"
	"slices"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrWorkspaceNotFound     = errors.New("workspace not found")
	ErrNoWorkspaceAvailable  = "no workspace available"
	ErrCouldNotSaveWorkspace = errors.New("could not save workspace")
)

type InMemWorkspaces struct {
	m                   sync.RWMutex
	workspaces          map[uuid.UUID]*idaccess_dom.Workspace
	memberships         map[uuid.UUID]*idaccess_dom.Membership
	isErrorExpected     bool
	isSaveErrorExpected bool
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
	imw.m.RLock()
	defer imw.m.RUnlock()

	w, ok := imw.workspaces[workspaceId]
	if !ok {
		return nil, fmt.Errorf("%w : %s", ErrWorkspaceNotFound, workspaceId)
	}
	for _, m := range imw.memberships {
		if m.IsSameWorkspace(w.Id()) && m.IsSameUser(userId) {
			return w.Clone(), nil
		}
	}
	return nil, fmt.Errorf("%w : %s", ErrWorkspaceNotFound, workspaceId)
}

func (imw *InMemWorkspaces) FindByName(ctx context.Context, userId uuid.UUID, name string) (*idaccess_dom.Workspace, error) {
	imw.m.RLock()
	defer imw.m.RUnlock()

	for _, w := range imw.workspaces {
		if w.Name() == name {
			for _, m := range imw.memberships {
				if m.WorkspaceId() == w.Id() && m.IsSameUser(userId) {
					return w.Clone(), nil
				}
			}
		}
	}

	return nil, fmt.Errorf("%w : %s", ErrWorkspaceNotFound, name)
}

func (imw *InMemWorkspaces) GetWorkspaces(ctx context.Context, userId uuid.UUID) ([]*idaccess_dom.Workspace, error) {
	imw.m.RLock()
	defer imw.m.RLock()

	wpIds := make([]uuid.UUID, 0, 10)
	for _, m := range imw.memberships {
		if m.IsSameUser(userId) {
			wpIds = append(wpIds, m.WorkspaceId())
		}
	}
	if len(wpIds) == 0 {
		return []*idaccess_dom.Workspace{}, nil
	}
	wp := make([]*idaccess_dom.Workspace, 0, 10)
	for _, w := range imw.workspaces {
		if slices.Contains(wpIds, w.Id()) {
			wp = append(wp, w)
		}
	}
	return wp, nil
}

func (imw *InMemWorkspaces) Save(ctx context.Context, userId uuid.UUID, workspace *idaccess_dom.Workspace) error {
	if imw.isSaveErrorExpected {
		return fmt.Errorf("%w : %s", ErrCouldNotSaveWorkspace, workspace.Id().String())
	}
	_, err := imw.FindById(ctx, userId, workspace.Id())
	if err != nil {
		return err
	}

	imw.m.RLock()
	defer imw.m.RLock()
	imw.workspaces[workspace.Id()] = workspace.Clone()
	return nil
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

func (imw *InMemWorkspaces) SetIsSaveErrorExpected(isSaveErrorExpected bool) {
	imw.isSaveErrorExpected = isSaveErrorExpected
}
