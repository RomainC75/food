package idaccess_infra

import (
	"errors"
	idaccess_dom "food/internal/modules/id-access/domain"
)

type InMemJWT struct {
	expectedToken           string
	isGenerateErrorExpected bool
	isVerifyError           bool
	verifyInfo              idaccess_dom.VerifyInfo
}

func NewInMemJWT() *InMemJWT {
	return &InMemJWT{}
}

func (imj *InMemJWT) Generate(user *idaccess_dom.User) (string, error) {
	if imj.isGenerateErrorExpected {
		return "", errors.New(idaccess_dom.ErrCouldNotGenerateJwt)
	}
	return imj.expectedToken, nil
}

func (imj *InMemJWT) Verify(tokenStr string) (idaccess_dom.VerifyInfo, error) {
	if imj.isVerifyError {
		return idaccess_dom.VerifyInfo{}, errors.New(idaccess_dom.ErrCouldNotVerifyToken)
	}
	return imj.verifyInfo, nil
}

func (imj *InMemJWT) SetupExpectedTokenError(expectedGenError bool) {
	imj.isGenerateErrorExpected = expectedGenError
}

func (imj *InMemJWT) SetupExpectedToken(expectedToken string) {
	imj.expectedToken = expectedToken
}
