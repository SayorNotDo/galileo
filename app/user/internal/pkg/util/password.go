package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return bytes, err
}

func ComparePassword(password string, hashed []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))
	return err == nil
}
