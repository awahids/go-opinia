package domain

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Fullname string
	Email    string
	Password string
	Phone    int
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
