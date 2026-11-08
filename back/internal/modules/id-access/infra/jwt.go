package idaccess_infra

import (
	"errors"
	"fmt"
	idaccess_dom "food/internal/modules/id-access/domain"
	shared_domain "food/internal/shared/domain"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidId = "invalid id"
)

type JWT struct {
	dateTime  shared_domain.DateTime
	secret    []byte
	expiresIn time.Duration
}

func NewJWT(dateTime shared_domain.DateTime, secret string, expiresIn time.Duration) *JWT {
	return &JWT{
		dateTime:  dateTime,
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
			ExpiresAt: jwt.NewNumericDate(j.dateTime.Now().Add(j.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(j.dateTime.Now()),
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
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithTimeFunc(j.dateTime.Now))
	if err != nil {
		log.Println("---> first : ", err.Error())
		return idaccess_dom.VerifyInfo{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return idaccess_dom.VerifyInfo{}, errors.New(fmt.Sprintf("%s : %s", ErrInvalidId, claims["user_id"].(string)))
		}

		return idaccess_dom.VerifyInfo{
			UserId:    userId,
			UserEmail: claims["user_email"].(string),
		}, nil
	}
	fmt.Println(err)
	return idaccess_dom.VerifyInfo{}, err
}
