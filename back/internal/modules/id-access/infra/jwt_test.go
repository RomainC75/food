package idaccess_infra

import (
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_infra "food/internal/shared/infra"
	utils_must "food/utils/must"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	refreshSecret    = "fake-secret"
	accessSecret     = "access-secret"
	refreshExpiresIn = time.Duration(time.Minute * 30)
	accessExpiresIn  = time.Duration(time.Minute * 15)

	userStringId = "33d89279-0e61-4972-95d1-579dbc2233d6"
	userId       = uuid.MustParse(userStringId)
	userEmail    = "fake@email.com"
	password     = "hashedPass"

	layout = "2006-01-02 15:04:05"

	generateNowStr = "2026-01-01 00:00:00"
	verifyNowStr   = "2026-01-01 00:15:00"
	generateNow    = utils_must.Must[time.Time](time.Parse(layout, generateNowStr))
	verifyNow      = utils_must.Must[time.Time](time.Parse(layout, verifyNowStr))
)

func TestJWTGenerateToken(t *testing.T) {
	t.Run("should generate token", func(t *testing.T) {
		dateTime := shared_infra.NewInMemDateTime().SetNow(generateNow)
		jwt := NewJWT(dateTime, refreshSecret, accessSecret, refreshExpiresIn, accessExpiresIn)

		user := idaccess_dom.NewUser(userId, userEmail, password)
		generatedTokens, err := jwt.Generate(user)

		assert.Nil(t, err)
		assert.NotEqual(t, "", generatedTokens.AccessToken)
		assert.NotEqual(t, "", generatedTokens.RefreshToken)

	})
}

func TestVerifyJwt(t *testing.T) {
	tests := []struct {
		name          string
		tokenToVerify string
		verifyNow     time.Time
		expectedError string
		reverseToken  bool
	}{
		{
			"should not verify outdated token",
			"",
			utils_must.Must[time.Time](time.Parse("2006-01-02 15:04:05", "2026-01-01 00:32:00")),
			"token is expired",
			false,
		},
		{
			"should not verify invalid token",
			"fake-token",
			verifyNow,
			"token is malformed",
			false,
		},
		{
			"should not verify if signature is invalid",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzM2Q4OTI3OS0wZTYxLTQ5NzItOTVkMS01NzlkYmMyMjMzZDYiLCJleHAiOjE3NjcyMjc0MDAsImlhdCI6MTc2NzIyNTYwMCwidXNlcl9pZCI6IjMzZDg5Mjc5LTBlNjEtNDk3Mi05NWQxLTU3OWRiYzIyMzNkNiJ9.5aeE6M3tgmb1Ry-EEtm3Py-D4gKOBwMEGG_TN6HkPKg",
			verifyNow,
			"signature is invalid",
			false,
		},
		{
			"should not verify if email is missing",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzM2Q4OTI3OS0wZTYxLTQ5NzItOTVkMS01NzlkYmMyMjMzZDYiLCJleHAiOjE3NjcyMjY1MDAsImlhdCI6MTc2NzIyNjUwMCwidXNlcl9pZCI6IjMzZDg5Mjc5LTBlNjEtNDk3Mi05NWQxLTU3OWRiYzIyMzNkNiJ9.Lm7VvYx2ymg2-FXvgfvHzeY0q1wcc9b0sV7oTQSZH74",
			verifyNow,
			"token has invalid claims",
			false,
		},
		{
			"should not verify if id is missing",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzM2Q4OTI3OS0wZTYxLTQ5NzItOTVkMS01NzlkYmMyMjMzZDYiLCJleHAiOjE3NjcyMjY1MDAsImlhdCI6MTc2NzIyNjUwMCwiZW1haWwiOiJmYWtlQGVtYWlsLmNvbSJ9.p34NxB-ccBPo4QOvrJ6prjgiqvB2AGGQKPjJG4yssBw",
			verifyNow,
			"token has invalid claims",
			false,
		},
		{
			"should not verify if id is invalid",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzM2Q4OTI3OS0wZTYxLTQ5NzItOTVkMS01NzlkYmMyMjMzZDYiLCJleHAiOjE3NjcyMjc0MDAsImlhdCI6MTc2NzIyNTYwMCwidXNlcl9pZCI6IjMzZDg5Mjc5LTBlNjEtNDk3Mi05NWQxLTU3OWRiYzIyMyIsInVzZXJfZW1haWwiOiJmYWtlQGVtYWlsLmNvbSJ9.V3J0NWUu6oot15k8J4Ap8Geenx_TXOPJIk7s9EGmSJA",
			verifyNow,
			ErrInvalidId,
			false,
		},
		{
			"should not verify an access token",
			"",
			verifyNow,
			"signature is invalid",
			true,
		},
		{
			"should verify",
			"",
			verifyNow,
			"",
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dateTime := shared_infra.NewInMemDateTime().SetNow(generateNow)
			jwt := NewJWT(dateTime, refreshSecret, accessSecret, refreshExpiresIn, accessExpiresIn)
			user := idaccess_dom.NewUser(userId, userEmail, password)

			// ? generate token
			generatedTokens, err := jwt.Generate(user)
			// ? test
			dateTime.SetNow(test.verifyNow)
			if test.tokenToVerify != "" {
				generatedTokens.RefreshToken = test.tokenToVerify
			}
			if test.reverseToken {
				refresh := generatedTokens.RefreshToken
				generatedTokens.RefreshToken = generatedTokens.AccessToken
				generatedTokens.AccessToken = refresh
			}

			infos, err := jwt.VerifyRefresh(generatedTokens.RefreshToken)
			if test.expectedError != "" {
				assert.Contains(t, err.Error(), test.expectedError)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, userId, infos.UserId)
		})
	}

}

func TestRefresh(t *testing.T) {
	dateTime := shared_infra.NewInMemDateTime().SetNow(generateNow)
	jwt := NewJWT(dateTime, refreshSecret, accessSecret, refreshExpiresIn, accessExpiresIn)
	user := idaccess_dom.NewUser(userId, userEmail, password)
	generatedTokens, _ := jwt.Generate(user)
	t.Run("should NOT refresh with a access token", func(t *testing.T) {
		_, err := jwt.Refresh(generatedTokens.AccessToken)
		assert.NotNil(t, err)
		assert.Contains(t, "invalid token", err.Error())
	})

	t.Run("should refresh with a refresh token", func(t *testing.T) {
		newRefreshToken, err := jwt.Refresh(generatedTokens.RefreshToken)
		assert.Nil(t, err)
		assert.NotEqual(t, generatedTokens.RefreshToken, newRefreshToken)
		assert.Contains(t, newRefreshToken, "eyJhbGc")
	})
}
