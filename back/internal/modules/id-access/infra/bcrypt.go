package idaccess_infra

import (
	"errors"
	"fmt"
	idaccess_dom "food/internal/modules/id-access/domain"
	"log"

	"golang.org/x/crypto/bcrypt"
)

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
	if err != nil {
		return "", errors.New(idaccess_dom.ErrHashingPassword)
	}
	log.Println("--> GENERATED ", password, string(hashed))
	return string(hashed), nil
}

func (bg *BcryptGen) IsValid(hash string, plain string) bool {
	log.Println("---> BCRYPT COMPARE /", hash, plain)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		fmt.Println("bcrypt", err.Error())
	}
	return err == nil
}
