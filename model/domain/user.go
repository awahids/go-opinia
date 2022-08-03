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

func (u *User) HashPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
