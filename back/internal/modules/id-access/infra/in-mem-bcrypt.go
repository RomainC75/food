package idaccess_infra

import (
	"errors"
	idaccess_dom "food/internal/modules/id-access/domain"
)

type InMemBcrypt struct {
	expectedHashed          string
	isGenerateErrorExpected bool
	isValid                 bool
}

func NewInMemBcrypt() *InMemBcrypt {
	return &InMemBcrypt{}
}

func (bg *InMemBcrypt) Generate(password string) (string, error) {
	if bg.isGenerateErrorExpected {
		return "", errors.New(idaccess_dom.ErrHashingPassword)
	}
	return bg.expectedHashed, nil
}
func (bg *InMemBcrypt) IsValid(hash string, password string) bool {
	return bg.isValid
}

func (bg *InMemBcrypt) SetNextHashPass(expectedHash string) {
	bg.expectedHashed = expectedHash
}

func (bg *InMemBcrypt) SetIsGenerateErrorExpected(isGenerateErrorExpected bool) {
	bg.isGenerateErrorExpected = isGenerateErrorExpected
}
