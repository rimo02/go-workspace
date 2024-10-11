package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err

}
