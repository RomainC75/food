package idaccess_infra

import "golang.org/x/crypto/bcrypt"

var (
	BCRYPT_COST = 10
)

type BcryptGen struct {
}

func NewBcrypt() *BcryptGen {
	return &BcryptGen{}
}

func (bg *BcryptGen) Generate(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), BCRYPT_COST)
	return string(hashed), err
}
func (bg *BcryptGen) IsValid(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
