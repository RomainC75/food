package idaccess_dom

import "errors"

var (
	ErrHashingPassword  = errors.New("error trying to hash the password")
	ErrNotValidPassword = errors.New("error password not valid")
)

type Bcrypt interface {
	Generate(password string) (string, error)
	IsValid(hash string, password string) bool
}
