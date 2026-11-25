package idaccess_dom

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCouldNotGenerateJwt = errors.New("could not generate jwt")
	ErrCouldNotVerifyToken = errors.New("could not verify token")
	ErrInvalidToken        = errors.New("invalid token")
)

type GeneratedTokens struct {
	AccessToken  string
	RefreshToken string
}
type VerifyInfo struct {
	UserId    uuid.UUID
	UserEmail string
}

type JWT interface {
	Generate(user *User) (GeneratedTokens, error)
	VerifyAccess(accessToken string) (VerifyInfo, error)
	Refresh(refreshToken string) (string, error)
}
