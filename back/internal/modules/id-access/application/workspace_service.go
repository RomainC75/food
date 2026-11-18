package idaccess_app

import (
	"context"
	"errors"
	"fmt"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"

	"github.com/google/uuid"
)

var (
	ErrTryingToCreateWorkspace = "error trying to create a workspace"
	ErrWorkspaceAlreadyExists  = errors.New("error workspace already exists")
)

type CreateWorkspaceCmd struct {
	Name   string
	UserId uuid.UUID
}

type WorkspaceSrv struct {
	uuidGen    shared_domain.Uuids
	workspaces idaccess_dom.Workspaces
	dateTime   shared_domain.DateTime
}

func NewWorkspaceSrv(
	uuidGen shared_domain.Uuids,
	workspaces idaccess_dom.Workspaces,
	dateTime shared_domain.DateTime,
) *WorkspaceSrv {
	return &WorkspaceSrv{
		uuidGen:    uuidGen,
		workspaces: workspaces,
		dateTime:   dateTime,
	}
}

func (wSrv *WorkspaceSrv) Create(ctx context.Context, cmd CreateWorkspaceCmd) (*idaccess_dom.Workspace, error) {
	_, err := wSrv.workspaces.FindByName(ctx, cmd.UserId, cmd.Name)
	if err == nil {
		return nil, fmt.Errorf("%w : %s", ErrWorkspaceAlreadyExists, cmd.Name)
	}

	workspaceId := wSrv.uuidGen.Generate()
	membershipId := wSrv.uuidGen.Generate()
	workspace := idaccess_dom.NewWorkspace(workspaceId, cmd.Name)
	membership := idaccess_dom.NewMembership(membershipId, cmd.UserId, workspaceId, idaccess_dom.Manager)

	err = wSrv.workspaces.CreateWithMembership(ctx, workspace, membership)
	if err != nil {
		return nil, errors.New(ErrTryingToCreateWorkspace)
	}

	return workspace, nil
}

func (wSrv *WorkspaceSrv) Get(ctx context.Context, userId uuid.UUID) ([]*idaccess_dom.Workspace, error) {
	return wSrv.workspaces.GetWorkspaces(ctx, userId)
}

func (wSrv *WorkspaceSrv) IsMember(
	ctx context.Context,
	userId uuid.UUID,
	workspaceID uuid.UUID,
) error {
	_, err := wSrv.workspaces.FindById(ctx, userId, workspaceID)
	return err
}
