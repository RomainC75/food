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
	user              = idaccess_dom.NewUser(userId, userEmail, userHashedPass)
	newUserReq        = dto_idaccess_req.SignupUser{
		Email:    userEmail,
		Password: userHashedPass,
	}
)

func TestSignupUser(t *testing.T) {
	// ? Next cases : password length / complexity ??
	type SignupUserSetup struct {
		existingUser              *idaccess_dom.User
		uuidExpected              uuid.UUID
		bcryptHashExpected        string
		isGenerateErrorExpected   bool
		isSavingUserErrorExpected bool
	}
	tests := []struct {
		name          string
		SignupUserReq dto_idaccess_req.SignupUser
		expectedError string
		setup         SignupUserSetup
	}{
		{
			"should not create duplicate user",
			newUserReq,
			ErrExistingUser,
			SignupUserSetup{
				user,
				uuid.New(),
				"",
				false,
				false,
			},
		},
		{
			"should not create duplicate user",
			newUserReq,
			idaccess_dom.ErrHashingPassword,
			SignupUserSetup{
				nil,
				uuid.New(),
				"",
				true,
				false,
			},
		},
		{
			"should not save user",
			newUserReq,
			idaccess_dom.ErrCouldNotSaveUser,
			SignupUserSetup{
				nil,
				uuid.New(),
				"",
				false,
				true,
			},
		},
		{
			"should save user",
			newUserReq,
			"",
			SignupUserSetup{
				nil,
				userId,
				"",
				false,
				false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			td := NewTestDriver(t).
				GivenUser(test.setup.existingUser).
				GivenUuid(userId).
				SetupBcryptGenResponse(test.setup.isGenerateErrorExpected, userHashedPass).
				SetupSavingUserError(test.setup.isSavingUserErrorExpected)

			createdUser, err := td.SignupUser(newUserReq)
			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
				return
			}
			assert.Nil(t, err)
			expectedUser := user
			assert.Nil(t, err)
			assert.Equal(t, expectedUser, createdUser)

			assert.Len(t, td.SavedUsers(), 1)
			assert.Equal(t, td.SavedUsers()[0], createdUser)
		})

	}
}

// wrong email
// wrong pass
// err generating token
// ok

func TestSignin(t *testing.T) {
	type SignupUserSetup struct {
		isBcryptPassValid  bool
		isGenTokenExpected bool
	}
	tests := []struct {
		name          string
		SignupUserReq dto_idaccess_req.SignupUser
		expectedToken string
		expectedError string
		setup         SignupUserSetup
	}{
		{
			name: "should not signup due to bad email",
			SignupUserReq: dto_idaccess_req.SigninUser{
				Email:    "badmail@email.com",
				Password: userClearPassword,
			},
			expectedToken: "",
			expectedError: ErrWrongEmailOrPassword,
			setup: SignupUserSetup{
				isBcryptPassValid:  true,
				isGenTokenExpected: false,
			},
		},
		{
			name: "should not signup due to bad pass",
			SignupUserReq: dto_idaccess_req.SigninUser{
				Email:    userEmail,
				Password: userClearPassword,
			},
			expectedToken: "",
			expectedError: ErrWrongEmailOrPassword,
			setup: SignupUserSetup{
				isBcryptPassValid:  false,
				isGenTokenExpected: false,
			},
		},
		{
			name: "should raise an error trying to generate a jwt",
			SignupUserReq: dto_idaccess_req.SigninUser{
				Email:    userEmail,
				Password: userClearPassword,
			},
			expectedToken: "",
			expectedError: idaccess_dom.ErrCouldNotGenerateJwt,
			setup: SignupUserSetup{
				isBcryptPassValid:  true,
				isGenTokenExpected: true,
			},
		},
		{
			name: "should sign the user in",
			SignupUserReq: dto_idaccess_req.SigninUser{
				Email:    userEmail,
				Password: userClearPassword,
			},
			expectedToken: "encoded-token",
			expectedError: "",
			setup: SignupUserSetup{
				isBcryptPassValid:  true,
				isGenTokenExpected: false,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			td := NewTestDriver(t).
				GivenUser(user).
				SetupExpectedTokenError(test.setup.isGenTokenExpected).
				SetBcryptIsValid(test.setup.isBcryptPassValid).
				SetupExpetedToken(test.expectedToken)

			tokenRes, err := td.SigninUser(test.SignupUserReq)

			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, test.expectedToken, tokenRes.Token)
			assert.Equal(t, test.SignupUserReq.Email, tokenRes.Email)
			assert.Equal(t, user.Id(), tokenRes.UserID)
		})
	}
}
