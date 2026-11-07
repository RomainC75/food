package idaccess_infra

import (
	idaccess_dom "food/internal/modules/id-access/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWT_KEY       = "super key"
	EXPIRE_IN_MIN = 30
)

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

type CustomClaim struct {
	jwt.RegisteredClaims
	UserId    string `json:"user_id"`
	UserEmail string `json:"user_email"`
}

func (j *JWT) Generate(user *idaccess_dom.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &CustomClaim{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(EXPIRE_IN_MIN))),
		},
		user.Id().String(),
		user.Email(),
	}
	return t.SignedString(JWT_KEY)

}

// func (j *JWT) Verify(tokenStr string) bool {
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
// 		return []byte(JWT_KEY), nil
// 	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
// 	if err != nil {
// 		return false
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		fmt.Println(claims["foo"], claims["nbf"])
// 		return true
// 	}
// 	fmt.Println(err)
// 	return false
// }
