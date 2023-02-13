package service

import (
	"Gurumu/features/autentikasi"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

type autentikasiUseCase struct {
	qry autentikasi.AutentikasiData
	vld *validator.Validate
}

func New(ad autentikasi.AutentikasiData) autentikasi.AutentikasiService {
	return &autentikasiUseCase{
		qry: ad,
		vld: validator.New(),
	}
}
func (aud *autentikasiUseCase) Login(email, password string) (string, autentikasi.Core, error) {
	res, err := aud.qry.Login(email)
	fmt.Println("ini :", err)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", autentikasi.Core{}, errors.New(msg)
	}

	if err := helper.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", autentikasi.Core{}, errors.New("password tidak sesuai ")
	}

	token, _ := helper.GenerateJWT(int(res.ID))

	return token, res, nil

}
