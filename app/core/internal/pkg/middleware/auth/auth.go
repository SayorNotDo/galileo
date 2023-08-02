package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	TokenExpiration = time.Hour * 24 * 30
)

type CustomClaims struct {
	ID          uint32
	Username    string
	AuthorityId int
	jwt.RegisteredClaims
}

type ReportClaims struct {
	Machine string `json:"machine,omitempty"`
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

func GenerateExecuteToken(r ReportClaims, key string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, r)
	signedString, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", errors.New("generate execute-token failed" + err.Error())
	}
	return signedString, nil
}
