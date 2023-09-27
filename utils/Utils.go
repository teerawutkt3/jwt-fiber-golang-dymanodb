package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetXRequestId() string {
	return uuid.New().String()
}

func FormatDate(t time.Time, patten string) string {
	return t.Format(patten)
}

func ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
