package idaccess_infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcrypt(t *testing.T) {
	pass := "my-password"
	bcrypt := NewBcrypt()
	hashed, err := bcrypt.Generate(pass)
	assert.Nil(t, err)
	assert.True(t, bcrypt.IsValid(hashed, pass))
}
