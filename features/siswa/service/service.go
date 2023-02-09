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
			msg = "data sudah terdaftar"
		} else {
			msg = "masalah pada server"
		}
		return siswa.Core{}, errors.New(msg)
	}

	return res, nil
}

func (suc *siswaUseCase) Profile(token interface{}) (siswa.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return siswa.Core{}, errors.New("id tidak valid")
	}

	res, err := suc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return siswa.Core{}, errors.New(msg)
	}
	return res, nil
}

func (suc *siswaUseCase) Update(token interface{}, updateData siswa.Core, avatar *multipart.FileHeader) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	if updateData.Email == "" {
		res, _ := suc.qry.Profile(uint(id))

		updateData.Email = res.Email
	}

	if avatar != nil {
		path, _ := helper.UploadStudentProfilePhotoS3(*avatar, updateData.Email)
		updateData.Avatar = path
	}

	// log.Println(updateData.Avatar)
	err := suc.qry.Update(uint(id), updateData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error update query: ", err.Error())
		return errors.New(msg)
	}
	return nil
}

func (suc *siswaUseCase) Delete(token interface{}) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	err := suc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}

	return nil
}
