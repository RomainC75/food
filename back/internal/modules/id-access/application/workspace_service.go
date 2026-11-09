package idaccess_app

import (
	"context"
	"errors"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"

	"github.com/google/uuid"
)

var (
	ErrTryingToCreateWorkspace = "error trying to create a workspace"
)

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

func (wSrv *WorkspaceSrv) Create(ctx context.Context, req dto_idaccess_req.CreateWorkspace) (*idaccess_dom.Workspace, error) {
	workspaceId := wSrv.uuidGen.Generate()
	membershipId := wSrv.uuidGen.Generate()

	userId := ctx.Value(idaccess_dom.TokenUserIdKey).(uuid.UUID)

	workspace := idaccess_dom.NewWorkspace(workspaceId, req.Name)
	membership := idaccess_dom.NewMembership(membershipId, userId, workspaceId, idaccess_dom.Manager)

	err := wSrv.workspaces.CreateWithMembership(ctx, workspace, membership)
	if err != nil {
		return nil, errors.New(ErrTryingToCreateWorkspace)
	}

	return workspace, nil
}
