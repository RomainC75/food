package idaccess_app

import (
	"context"
	"errors"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"
)

type IdAccessSrv struct {
	uuidGen shared_domain.Uuids
	bcrypt  idaccess_dom.Bcrypt
	users   idaccess_dom.Users
}

func NewIdAccessSrv(uuidGen shared_domain.Uuids, bcryptHandler idaccess_dom.Bcrypt, users idaccess_dom.Users) *IdAccessSrv {
	return &IdAccessSrv{
		uuidGen: uuidGen,
		bcrypt:  bcryptHandler,
		users:   users,
	}
}

func (iaSrv *IdAccessSrv) CreateUser(ctx context.Context, newUserReq dto_idaccess_req.CreateUser) (*idaccess_dom.User, error) {
	_, err := iaSrv.users.FindByEmail(ctx, newUserReq.Email)
	if err == nil {
		return nil, err
	}
	newUserId := iaSrv.uuidGen.Generate()
	hashedPass, err := iaSrv.bcrypt.Generate(newUserReq.Email)
	if err != nil {
		return nil, errors.New(idaccess_dom.ErrGeneratingBcrypt)
	}

	newUser := idaccess_dom.NewUser(newUserId, newUserReq.Email, hashedPass)
	err = iaSrv.users.Create(ctx, newUser)
	return newUser, err

}
