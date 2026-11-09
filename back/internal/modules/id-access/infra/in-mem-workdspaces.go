package idaccess_infra

import (
	"context"
	"errors"
	idaccess_dom "food/internal/modules/id-access/domain"
)

type InMemWorkspaces struct {
	workspaces      []*idaccess_dom.Workspace
	memberships     []*idaccess_dom.Membership
	isErrorExpected bool
}

func NewInMemWorkspaces() *InMemWorkspaces {
	return &InMemWorkspaces{
		workspaces:  make([]*idaccess_dom.Workspace, 0, 2),
		memberships: make([]*idaccess_dom.Membership, 0, 2),
	}
}

func (imw *InMemWorkspaces) CreateWithMembership(ctx context.Context, workspace *idaccess_dom.Workspace, membership *idaccess_dom.Membership) error {
	if imw.isErrorExpected {
		return errors.New(idaccess_dom.ErrCouldNotCreateWorkspace)
	}
	imw.workspaces = append(imw.workspaces, workspace.Clone())
	imw.memberships = append(imw.memberships, membership.Clone())
	return nil
}

// ===== TEST ===

func (imw *InMemWorkspaces) SavedMemberships() []*idaccess_dom.Membership {
	return imw.memberships
}

func (imw *InMemWorkspaces) SavedWorkspaces() []*idaccess_dom.Workspace {
	return imw.workspaces
}

func (imw *InMemWorkspaces) SetIsErrorExpected(isErrorExpected bool) {
	imw.isErrorExpected = isErrorExpected
}
