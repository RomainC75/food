package idaccess_infra

type InMemBcrypt struct {
	expectedHashed string
	isValid        bool
}

func NewInMemBcrypt() *InMemBcrypt {
	return &InMemBcrypt{}
}

func (bg *InMemBcrypt) Generate(password string) (string, error) {
	return bg.expectedHashed, nil
}
func (bg *InMemBcrypt) IsValid(hash string, password string) bool {
	return bg.isValid
}

func (bg *InMemBcrypt) SetNextHashPass(expectedHash string) {
	bg.expectedHashed = expectedHash
}
