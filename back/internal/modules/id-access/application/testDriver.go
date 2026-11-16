package idaccess_app

import (
	"context"
	idaccess_dom "food/internal/modules/id-access/domain"
	idaccess_infra "food/internal/modules/id-access/infra"
	platform_http "food/internal/platform/http"
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

func (td *TestDriver) SignupUser(cmd SignupCmd) (*idaccess_dom.User, error) {
	return td.userSrv.SignupUser(td.ctx, cmd)
}

func (td *TestDriver) SigninUser(cmd SigninCmd) (SigninResult, error) {
	return td.userSrv.SigninUser(td.ctx, cmd)
}

func (td *TestDriver) CreateWorkspace(createWCmd CreateWorkspaceCmd) (*idaccess_dom.Workspace, error) {
	return td.workspaceSrv.Create(td.ctx, createWCmd)
}

func (td *TestDriver) GetWorkspaces(userId uuid.UUID) ([]*idaccess_dom.Workspace, error) {
	return td.workspaceSrv.Get(td.ctx, userId)
}

// ======== SETUP =========

func (td *TestDriver) GivenUserIdContext(userId uuid.UUID) *TestDriver {
	td.ctx = platform_http.WithUserId(td.ctx, userId)
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

func (td *TestDriver) SetupExpectedToken(expectedToken string) *TestDriver {
	td.jwtGen.SetupExpectedToken(expectedToken)
	return td
}

func (td *TestDriver) SetupExpectedWorkspaceError(isErrorExpected bool) {
	td.workspaces.SetIsErrorExpected(isErrorExpected)
}

func (td *TestDriver) SavedUsers() map[uuid.UUID]*idaccess_dom.User {
	return td.users.SavedUsers()
}

func (td *TestDriver) SavedMemberships() map[uuid.UUID]*idaccess_dom.Membership {
	return td.workspaces.SavedMemberships()
}

func (td *TestDriver) SavedWorkspaces() map[uuid.UUID]*idaccess_dom.Workspace {
	return td.workspaces.SavedWorkspaces()
}
