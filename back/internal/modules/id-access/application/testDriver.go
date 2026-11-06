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
	users       idaccess_dom.Users
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

func (td *TestDriver) CreateUser(newUserReq dto_idaccess_req.CreateUser) (*idaccess_dom.User, error) {
	return td.IdAccessSrv.CreateUser(td.ctx, newUserReq)
}

func (td *TestDriver) GivenHashedPass(expectedHash string) *TestDriver {
	td.bcrypt.SetNextHashPass(expectedHash)
	return td
}

func (td *TestDriver) GivenUuid(id uuid.UUID) *TestDriver {
	td.uuidGen.SetNext(id)
	return td
}
