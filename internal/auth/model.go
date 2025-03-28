package auth

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Phone     string
	CreatedAt time.Time
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
