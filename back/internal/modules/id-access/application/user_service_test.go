package idaccess_app

import (
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
	signupCmd         = SignupCmd{
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
		signupCmd     SignupCmd
		expectedError string
		setup         SignupUserSetup
	}{
		{
			"should not create duplicate user",
			signupCmd,
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
			signupCmd,
			idaccess_dom.ErrHashingPassword.Error(),
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
			signupCmd,
			idaccess_dom.ErrCouldNotSaveUser.Error(),
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
			signupCmd,
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

			createdUser, err := td.SignupUser(test.signupCmd)
			if test.expectedError != "" {
				assert.EqualError(t, err, test.expectedError)
				return
			}
			assert.Nil(t, err)
			expectedUser := user
			assert.Nil(t, err)
			assert.Equal(t, expectedUser, createdUser)

			assert.Len(t, td.SavedUsers(), 1)
			assert.Equal(t, td.SavedUsers()[test.setup.uuidExpected], createdUser)
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
		SignupUserReq SigninCmd
		expectedToken string
		expectedError string
		setup         SignupUserSetup
	}{
		{
			name: "should not signup due to bad email",
			SignupUserReq: SigninCmd{
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
			SignupUserReq: SigninCmd{
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
			SignupUserReq: SigninCmd{
				Email:    userEmail,
				Password: userClearPassword,
			},
			expectedToken: "",
			expectedError: idaccess_dom.ErrCouldNotGenerateJwt.Error(),
			setup: SignupUserSetup{
				isBcryptPassValid:  true,
				isGenTokenExpected: true,
			},
		},
		{
			name: "should sign the user in",
			SignupUserReq: SigninCmd{
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
				SetupExpectedToken(test.expectedToken)

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
