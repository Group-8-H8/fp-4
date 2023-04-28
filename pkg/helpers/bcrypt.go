package helpers

import (
	"github.com/fydhfzh/fp-4/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, errs.Errs) {
	seed := 8

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), seed)

	if err != nil {
		return "", errs.NewBadRequestError(err.Error())
	}

	return string(passwordBytes), nil
}

func ComparePassword(hashed string, password string) (bool) {
	bytesHashed := []byte(hashed)
	bytesPassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(bytesHashed, bytesPassword)

	return err == nil
}