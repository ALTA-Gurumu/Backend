package helper

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return "", errors.New("password process error")
	}

	return string(hashed), nil
}

func CheckPassword(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		log.Println("login compare", err.Error())
		return errors.New("password tidak sesuai ")
	}
	return nil
}
