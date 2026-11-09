package idaccess_app

import (
	"context"
	"errors"
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_app_mod "food/internal/modules/id-access/application/models"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"
)

var (
	ErrExistingUser         = "error : user exists"
	ErrWrongEmailOrPassword = "wrong email or password"
	ErrDecodingToken        = "error trying to decode the token"
)

type UserSrv struct {
	uuidGen shared_domain.Uuids
	bcrypt  idaccess_dom.Bcrypt
	jwtGen  idaccess_dom.JWT
	users   idaccess_dom.Users
}

func NewUserSrv(
	uuidGen shared_domain.Uuids,
	bcryptHandler idaccess_dom.Bcrypt,
	jwtGen idaccess_dom.JWT,
	users idaccess_dom.Users,
) *UserSrv {
	return &UserSrv{
		uuidGen: uuidGen,
		bcrypt:  bcryptHandler,
		jwtGen:  jwtGen,
		users:   users,
	}
}

func (uSrv *UserSrv) SignupUser(ctx context.Context, newUserReq dto_idaccess_req.SignupUser) (*idaccess_dom.User, error) {
	_, err := uSrv.users.FindByEmail(ctx, newUserReq.Email)

	if err == nil {
		return nil, errors.New(ErrExistingUser)
	}
	newUserId := uSrv.uuidGen.Generate()
	hashedPass, err := uSrv.bcrypt.Generate(newUserReq.Password)
	if err != nil {
		return nil, err
	}

	newUser := idaccess_dom.NewUser(newUserId, newUserReq.Email, hashedPass)
	err = uSrv.users.Create(ctx, newUser)
	return newUser, err
}

func (uSrv *UserSrv) SigninUser(ctx context.Context, signinUserReq dto_idaccess_req.SigninUser) (idaccess_app_mod.SigninResult, error) {
	foundUser, err := uSrv.users.FindByEmail(ctx, signinUserReq.Email)

	if err != nil {
		return idaccess_app_mod.SigninResult{}, errors.New(ErrWrongEmailOrPassword)
	}

	isPassValid := uSrv.bcrypt.IsValid(foundUser.Password(), signinUserReq.Password)
	if !isPassValid {
		return idaccess_app_mod.SigninResult{}, errors.New(ErrWrongEmailOrPassword)
	}

	tokenStr, err := uSrv.jwtGen.Generate(foundUser)
	return idaccess_app_mod.SigninResult{
		UserID: foundUser.Id(),
		Email:  foundUser.Email(),
		Token:  tokenStr,
	}, err
}

func (uSrv *UserSrv) Verify(tokenStr string) (idaccess_dom.VerifyInfo, error) {

	data, err := uSrv.jwtGen.Verify(tokenStr)
	if err != nil {
		return data, errors.New(ErrDecodingToken)
	}
	return data, nil
}
