package idaccess_app

import (
	dto_idaccess_req "food/internal/api/dtos/id-access/req"
	idaccess_dom "food/internal/modules/id-access/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	userId            = uuid.MustParse("1d3cd2bf-7f9e-4e0c-8d38-2ae227e234c2")
	userEmail         = "name@email.com"
	userClearPassword = "clear-pass"
	userHashedPass    = "hashed pass"
)

func TestCreateUser(t *testing.T) {
	td := NewTestDriver(t).GivenHashedPass(userHashedPass).GivenUuid(userId)
	newUserReq := dto_idaccess_req.CreateUser{
		Email:    userEmail,
		Password: userHashedPass,
	}

	createdUser, err := td.CreateUser(newUserReq)
	expectedUser := idaccess_dom.NewUser(userId, userEmail, userHashedPass)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, createdUser)
}
