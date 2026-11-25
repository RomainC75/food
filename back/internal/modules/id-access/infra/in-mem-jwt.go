package idaccess_infra

import (
	idaccess_dom "food/internal/modules/id-access/domain"
)

type InMemJWT struct {
	expectedAccessToken     string
	expectedRefreshToken    string
	expectedNewAccessToken  string
	isGenerateErrorExpected bool
	isVerifyError           bool
	verifyInfo              idaccess_dom.VerifyInfo
}

func NewInMemJWT() *InMemJWT {
	return &InMemJWT{}
}

func (imj *InMemJWT) Generate(user *idaccess_dom.User) (idaccess_dom.GeneratedTokens, error) {
	if imj.isGenerateErrorExpected {
		return idaccess_dom.GeneratedTokens{}, idaccess_dom.ErrCouldNotGenerateJwt
	}
	return idaccess_dom.GeneratedTokens{
		AccessToken:  imj.expectedAccessToken,
		RefreshToken: imj.expectedRefreshToken,
	}, nil
}

func (imj *InMemJWT) VerifyAccess(accessToken string) (idaccess_dom.VerifyInfo, error) {
	if imj.isVerifyError {
		return idaccess_dom.VerifyInfo{}, idaccess_dom.ErrCouldNotVerifyToken
	}
	return imj.verifyInfo, nil
}

func (imj *InMemJWT) Refresh(refreshToken string) (string, error) {
	return imj.expectedNewAccessToken, nil
}

func (imj *InMemJWT) SetupExpectedTokenError(expectedGenError bool) {
	imj.isGenerateErrorExpected = expectedGenError
}

func (imj *InMemJWT) SetupExpectedTokens(expectedRefreshToken string, expectedAccessToken string) {
	imj.expectedRefreshToken = expectedRefreshToken
	imj.expectedAccessToken = expectedAccessToken
}

func (imj *InMemJWT) SetupNewAccessToken(expectedNewAccessToken string) {
	imj.expectedNewAccessToken = expectedNewAccessToken

}
