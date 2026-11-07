package idaccess_app

import (
	"context"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_dom "food/internal/modules/id-access/domain"
	idaccess_infra "food/internal/modules/id-access/infra"
	shared_infra "food/internal/shared/infra"
	"testing"

	"github.com/google/uuid"
)

type TestDriver struct {
	t           *testing.T
	ctx         context.Context
	users       *idaccess_infra.InMemoryUsers
	uuidGen     *shared_infra.InMemUuidGen
	bcrypt      *idaccess_infra.InMemBcrypt
	IdAccessSrv *IdAccessSrv
}

func NewTestDriver(t *testing.T) *TestDriver {
	users := idaccess_infra.NewInMemoryUsers()
	uuidGen := shared_infra.NewInMemUuidGen()
	bcrypt := idaccess_infra.NewInMemBcrypt()
	return &TestDriver{
		t:           t,
		ctx:         context.Background(),
		users:       users,
		uuidGen:     uuidGen,
		bcrypt:      bcrypt,
		IdAccessSrv: NewIdAccessSrv(uuidGen, bcrypt, users),
	}
}

func (td *TestDriver) SignupUser(newUserReq dto_idaccess_req.SignupUser) (*idaccess_dom.User, error) {
	return td.IdAccessSrv.SignupUser(td.ctx, newUserReq)
}

// ======== SETUP =========

func (td *TestDriver) SetupBcryptResponse(isGenerateError bool, expectedHash string) *TestDriver {
	td.bcrypt.SetIsGenerateErrorExpected(isGenerateError)
	td.bcrypt.SetNextHashPass(expectedHash)
	return td
}

func (td *TestDriver) GivenUuid(id uuid.UUID) *TestDriver {
	td.uuidGen.SetNext(id)
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

func (td *TestDriver) SavedUsers() []*idaccess_dom.User {
	return td.users.SavedUsers()
}
