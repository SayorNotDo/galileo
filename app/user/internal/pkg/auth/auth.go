package auth

import (
	"errors"
	"fmt"
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
		return "", errors.New(fmt.Sprintf("generate token failed: %s", err.Error()))
	}
	return signedString, nil
}
