package idaccess_app

import (
	"context"
	"errors"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"
	"food/utils/displayer"

	"github.com/google/uuid"
)

var (
	ErrExistingUser         = "error : user exists"
	ErrWrongEmailOrPassword = "wrong email or password"
	ErrDecodingToken        = "error trying to decode the token"
)

type Credentials struct {
	Email    string
	Password string
}

type SignupCmd = Credentials
type LoginCmd = Credentials

type LoginResult struct {
	UserID       uuid.UUID
	Email        string
	AccessToken  string
	RefreshToken string
}

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

func (uSrv *UserSrv) SignupUser(ctx context.Context, cmd SignupCmd) (*idaccess_dom.User, error) {
	_, err := uSrv.users.FindByEmail(ctx, cmd.Email)

	if err == nil {
		return nil, errors.New(ErrExistingUser)
	}
	newUserId := uSrv.uuidGen.Generate()
	hashedPass, err := uSrv.bcrypt.Generate(cmd.Password)
	if err != nil {
		return nil, err
	}

	newUser := idaccess_dom.NewUser(newUserId, cmd.Email, hashedPass)
	err = uSrv.users.Create(ctx, newUser)
	return newUser, err
}

func (uSrv *UserSrv) LoginUser(ctx context.Context, cmd LoginCmd) (LoginResult, error) {
	foundUser, err := uSrv.users.FindByEmail(ctx, cmd.Email)

	if err != nil {
		return LoginResult{}, errors.New(ErrWrongEmailOrPassword)
	}

	isPassValid := uSrv.bcrypt.IsValid(foundUser.Password(), cmd.Password)
	if !isPassValid {
		return LoginResult{}, errors.New(ErrWrongEmailOrPassword)
	}

	generatedTokens, err := uSrv.jwtGen.Generate(foundUser)
	displayer.PrettyDisplay("tokens : ", generatedTokens)
	return LoginResult{
		UserID:       foundUser.Id(),
		Email:        foundUser.Email(),
		AccessToken:  generatedTokens.AccessToken,
		RefreshToken: generatedTokens.RefreshToken,
	}, err
}

func (uSrv *UserSrv) Verify(accessToken string) (idaccess_dom.VerifyInfo, error) {
	data, err := uSrv.jwtGen.VerifyAccess(accessToken)
	if err != nil {
		return data, errors.New(ErrDecodingToken)
	}
	return data, nil
}

func (uSrv *UserSrv) Refresh(refreshToken string) (string, error) {

	accessToken, err := uSrv.jwtGen.Refresh(refreshToken)
	if err != nil {
		return accessToken, errors.New(ErrDecodingToken)
	}
	return accessToken, nil
}
