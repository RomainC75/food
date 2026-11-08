package idaccess_infra

import (
	"fmt"
	idaccess_dom "food/internal/modules/id-access/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type JWT struct {
	secret    []byte
	expiresIn time.Duration
}

func NewJWT(secret string, expiresIn time.Duration) *JWT {
	return &JWT{
		secret:    []byte(secret),
		expiresIn: expiresIn,
	}
}

type CustomClaim struct {
	jwt.RegisteredClaims
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
}

func (j *JWT) Generate(user *idaccess_dom.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user.Id().String(),
		},
		UserID:    user.Id().String(),
		UserEmail: user.Email(),
	})

	signedString, err := token.SignedString(j.secret)

	if err != nil {
		logrus.Warn("JWT Generate() : ", err.Error())
		return "", fmt.Errorf("error creating jwt string: %v", err)
	}

	return signedString, nil
}

func (j *JWT) Verify(tokenStr string) (idaccess_dom.VerifyInfo, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(j.secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return idaccess_dom.VerifyInfo{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims, claims["user_id"], claims["user_email"])
		userId, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return idaccess_dom.VerifyInfo{}, err
		}
		return idaccess_dom.VerifyInfo{
			UserId:    userId,
			UserEmail: claims["user_email"].(string),
		}, nil
	}
	fmt.Println(err)
	return idaccess_dom.VerifyInfo{}, err
}
