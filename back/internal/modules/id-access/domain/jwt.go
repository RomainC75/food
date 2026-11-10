package idaccess_dom

import "github.com/google/uuid"

var (
	ErrCouldNotGenerateJwt = "could not generate jwt"
	ErrCouldNotVerifyToken = "could not verify token"
)

type VerifyInfo struct {
	UserId    uuid.UUID
	UserEmail string
}

type JWT interface {
	Generate(user *User) (string, error)
	Verify(token string) (VerifyInfo, error)
}
