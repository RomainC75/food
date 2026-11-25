package idaccess_infra

import (
	"errors"
	"fmt"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidId = "invalid id"
)

type JWT struct {
	dateTime         shared_domain.DateTime
	refreshSecret    string
	accessSecret     string
	refreshExpiresIn time.Duration
	accessExpiresIn  time.Duration
}

func NewJWT(dateTime shared_domain.DateTime, refreshsecret string, accessSecret string, refreshExpiresIn time.Duration, accessExpiresIn time.Duration) *JWT {
	return &JWT{
		dateTime:         dateTime,
		refreshSecret:    refreshsecret,
		accessSecret:     accessSecret,
		refreshExpiresIn: refreshExpiresIn,
		accessExpiresIn:  accessExpiresIn,
	}
}

type CustomClaim struct {
	jwt.RegisteredClaims
	UserID    string `json:"user_id"`
	UserEmail string `json:"email"`
}

func (j *JWT) Generate(user *idaccess_dom.User) (idaccess_dom.GeneratedTokens, error) {
	refreshToken, err := j.generateString(j.refreshSecret, user.Id(), user.Email(), j.refreshExpiresIn)
	if err != nil {
		logrus.Warn("JWT Generate() : ", err.Error())
		return idaccess_dom.GeneratedTokens{}, fmt.Errorf("error creating jwt string: %v", err)
	}
	accessToken, err := j.generateString(j.accessSecret, user.Id(), user.Email(), j.accessExpiresIn)
	if err != nil {
		logrus.Warn("JWT Generate() : ", err.Error())
		return idaccess_dom.GeneratedTokens{}, fmt.Errorf("error creating jwt string: %v", err)
	}

	return idaccess_dom.GeneratedTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (j *JWT) VerifyAccess(accessToken string) (idaccess_dom.VerifyInfo, error) {
	return j.verify(accessToken, j.accessSecret)
}

func (j *JWT) VerifyRefresh(refreshToken string) (idaccess_dom.VerifyInfo, error) {
	return j.verify(refreshToken, j.refreshSecret)
}

func (j *JWT) verify(tokenStr string, secret string) (idaccess_dom.VerifyInfo, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithTimeFunc(j.dateTime.Now))
	if err != nil {
		return idaccess_dom.VerifyInfo{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, ok := claims["user_id"]
		if !ok {
			return idaccess_dom.VerifyInfo{}, errors.New("could not get user_id from token")
		}
		userUuid, err := uuid.Parse(userId.(string))
		if err != nil {
			return idaccess_dom.VerifyInfo{}, errors.New(fmt.Sprintf("%s : %s", ErrInvalidId, claims["user_id"].(string)))
		}

		return idaccess_dom.VerifyInfo{
			UserId:    userUuid,
			UserEmail: claims["email"].(string),
		}, nil
	}
	fmt.Println(err)
	return idaccess_dom.VerifyInfo{}, err
}

func (j *JWT) Refresh(refreshToken string) (string, error) {
	userInfo, err := j.VerifyRefresh(refreshToken)
	if err != nil {
		return "", idaccess_dom.ErrInvalidToken
	}
	return j.generateString(string(j.accessSecret), userInfo.UserId, userInfo.UserEmail, j.accessExpiresIn)
}

func (j *JWT) generateString(secret string, userId uuid.UUID, email string, expiresIn time.Duration) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(j.dateTime.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(j.dateTime.Now()),
			Subject:   userId.String(),
		},
		UserID:    userId.String(),
		UserEmail: email,
	})
	return refreshToken.SignedString([]byte(secret))
}
