package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func HashPassword(password string) (string, error) {
	if len(strings.TrimSpace(password)) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
