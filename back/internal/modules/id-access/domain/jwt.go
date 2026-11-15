package idaccess_dom

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCouldNotGenerateJwt = errors.New("could not generate jwt")
	ErrCouldNotVerifyToken = errors.New("could not verify token")
)

type VerifyInfo struct {
	UserId    uuid.UUID
	UserEmail string
}

type JWT interface {
	Generate(user *User) (string, error)
	Verify(token string) (VerifyInfo, error)
}
