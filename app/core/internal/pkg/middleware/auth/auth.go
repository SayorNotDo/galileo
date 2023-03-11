package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	ID          uint32
	Username    string
	AuthorityId int
	jwt.RegisteredClaims
}

func CreateToken(c CustomClaims, key string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate token failed" + err.Error())
	}
	return signedString, nil
}
