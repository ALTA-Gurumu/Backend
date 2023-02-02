package service

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator"
)

type guruUseCase struct {
	qry guru.GuruData
	vld *validator.Validate
}

func New(gd guru.GuruData) guru.GuruService {
	return &guruUseCase{
		qry: gd,
		vld: validator.New(),
	}
}

// Register implements guru.GuruService
func (guc *guruUseCase) Register(newGuru guru.Core) (guru.Core, error) {
	hashed, err := helper.GeneratePassword(newGuru.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return guru.Core{}, errors.New("password process error")
	}

	err = guc.vld.Struct(&newGuru)
	if err != nil {
		msg := helper.ValidationErrorHandle(err)
		fmt.Println("msg", msg)
		return guru.Core{}, errors.New(msg)
	}

	newGuru.Password = string(hashed)
	res, err := guc.qry.Register(newGuru)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return guru.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements guru.GuruService
func (*guruUseCase) Delete(token interface{}) error {
	panic("unimplemented")
}

// Profile implements guru.GuruService
func (*guruUseCase) Profile(token interface{}) (guru.Core, error) {
	panic("unimplemented")
}

// Update implements guru.GuruService
func (*guruUseCase) Update(token interface{}, updateData guru.Core, avatar *multipart.FileHeader) (guru.Core, error) {
	panic("unimplemented")
}