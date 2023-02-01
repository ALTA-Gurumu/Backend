package service

import (
	"Gurumu/features/siswa"
	"Gurumu/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
)

type siswaUseCase struct {
	qry siswa.SiswaData
	vld *validator.Validate
}

func New(sd siswa.SiswaData) siswa.SiswaService {
	return &siswaUseCase{
		qry: sd,
		vld: validator.New(),
	}
}

func (suc *siswaUseCase) Register(newStudent siswa.Core) (siswa.Core, error) {
	err := suc.vld.Struct(newStudent)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return siswa.Core{}, errors.New("validation error")
	}

	hashed, err := helper.GeneratePassword(newStudent.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return siswa.Core{}, errors.New("password process error")
	}

	newStudent.Password = hashed

	res, err := suc.qry.Register(newStudent)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data already exist"
		} else {
			msg = "server problem"
		}
		return siswa.Core{}, errors.New(msg)
	}

	return res, nil
}

func (suc *siswaUseCase) Profile(token interface{}) (siswa.Core, error) {
	return siswa.Core{}, nil
}

func (suc *siswaUseCase) Update(token interface{}, updateData siswa.Core, avatar *multipart.FileHeader) (siswa.Core, error) {
	return siswa.Core{}, nil
}

func (suc *siswaUseCase) Delete(token interface{}) error {
	return nil
}