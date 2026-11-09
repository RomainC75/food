package idaccess_app

import (
	"context"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_app_mod "food/internal/modules/id-access/application/models"
	idaccess_dom "food/internal/modules/id-access/domain"
	idaccess_infra "food/internal/modules/id-access/infra"
	shared_infra "food/internal/shared/infra"
	"testing"

	"github.com/google/uuid"
)

type TestDriver struct {
	t            *testing.T
	ctx          context.Context
	users        *idaccess_infra.InMemoryUsers
	workspaces   *idaccess_infra.InMemWorkspaces
	uuidGen      *shared_infra.InMemUuidGen
	bcrypt       *idaccess_infra.InMemBcrypt
	jwtGen       *idaccess_infra.InMemJWT
	userSrv      *UserSrv
	workspaceSrv *WorkspaceSrv
}

func NewTestDriver(t *testing.T) *TestDriver {
	users := idaccess_infra.NewInMemoryUsers()
	workspaces := idaccess_infra.NewInMemWorkspaces()
	uuidGen := shared_infra.NewInMemUuidGen()
	bcrypt := idaccess_infra.NewInMemBcrypt()
	jwtGen := idaccess_infra.NewInMemJWT()
	dateTime := shared_infra.NewInMemDateTime()
	return &TestDriver{
		t:            t,
		ctx:          context.Background(),
		users:        users,
		workspaces:   workspaces,
		uuidGen:      uuidGen,
		bcrypt:       bcrypt,
		jwtGen:       jwtGen,
		userSrv:      NewUserSrv(uuidGen, bcrypt, jwtGen, users),
		workspaceSrv: NewWorkspaceSrv(uuidGen, workspaces, dateTime),
	}
}

func (td *TestDriver) SignupUser(newUserReq dto_idaccess_req.SignupUser) (*idaccess_dom.User, error) {
	return td.userSrv.SignupUser(td.ctx, newUserReq)
}

func (td *TestDriver) SigninUser(newUserReq dto_idaccess_req.SignupUser) (idaccess_app_mod.SigninResult, error) {
	return td.userSrv.SigninUser(td.ctx, newUserReq)
}

func (td *TestDriver) CreateWorkspace(ws dto_idaccess_req.CreateWorkspace) (*idaccess_dom.Workspace, error) {
	return td.workspaceSrv.Create(td.ctx, ws)
}

// ======== SETUP =========

func (td *TestDriver) GivenUserIdContext(userId uuid.UUID) *TestDriver {
	td.ctx = context.WithValue(td.ctx, idaccess_dom.TokenUserIdKey, userId)
	return td
}

func (td *TestDriver) SetupBcryptGenResponse(isGenerateError bool, expectedHash string) *TestDriver {
	td.bcrypt.SetisErrorExpected(isGenerateError)
	td.bcrypt.SetNextHashPass(expectedHash)
	return td
}

func (td *TestDriver) SetupBcryptError(isErrorExpected bool) *TestDriver {
	td.bcrypt.SetisErrorExpected(isErrorExpected)
	return td
}

func (td *TestDriver) SetBcryptIsValid(isValid bool) *TestDriver {
	td.bcrypt.SetIsValid(isValid)
	return td
}

func (td *TestDriver) GivenUuid(id uuid.UUID) *TestDriver {
	td.uuidGen.SetNext(id)
	return td
}

func (td *TestDriver) GivenUuids(ids ...uuid.UUID) *TestDriver {
	td.uuidGen.SetMultipleUuids(ids)
	return td
}

func (td *TestDriver) GivenUser(existingUser *idaccess_dom.User) *TestDriver {
	if existingUser != nil {
		td.users.Create(td.ctx, existingUser)
	}
	return td
}

func (td *TestDriver) SetupSavingUserError(isSavingUserErrorExpected bool) *TestDriver {
	td.users.SetupIsSavingUserError(isSavingUserErrorExpected)
	return td
}

func (td *TestDriver) SetupExpectedTokenError(expectGenTokenError bool) *TestDriver {
	td.jwtGen.SetupExpectedTokenError(expectGenTokenError)
	return td
}

func (td *TestDriver) SetupExpetedToken(expectedToken string) *TestDriver {
	td.jwtGen.SetupExpectedToken(expectedToken)
	return td
}

func (td *TestDriver) SetupExpectedWorkspaceError(isErrorExpected bool) {
	td.workspaces.SetIsErrorExpected(isErrorExpected)
}

func (td *TestDriver) SavedUsers() []*idaccess_dom.User {
	return td.users.SavedUsers()
}

func (td *TestDriver) SavedMemberships() []*idaccess_dom.Membership {
	return td.workspaces.SavedMemberships()
}

func (td *TestDriver) SavedWorkspaces() []*idaccess_dom.Workspace {
	return td.workspaces.SavedWorkspaces()
}
