package password

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct{}

func (p Password) Hashed(password string) string {
	var passwordByte = []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func (p Password) CompareHashedAndPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
