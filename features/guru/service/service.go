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
func (guc *guruUseCase) Delete(token interface{}) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("data not found")
	}

	err := guc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return errors.New(msg)
	}

	return nil
}

// Profile implements guru.GuruService
func (*guruUseCase) Profile(token interface{}) (guru.Core, error) {
	panic("unimplemented")
}

// Update implements guru.GuruService
func (guc *guruUseCase) Update(token interface{}, updateData guru.Core, avatar *multipart.FileHeader, ijazah *multipart.FileHeader) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("token tidak valid")
	}

	if err := guc.vld.Struct(&updateData); err != nil {
		log.Println(err)
		msg := helper.ValidationErrorHandle(err)
		return errors.New(msg)
	}
	if avatar == nil {
		if err := guc.qry.Update(uint(userID), updateData); err != nil {
			log.Println(err)
			msg := ""
			if strings.Contains(err.Error(), "tidak ditemukan") {
				msg = err.Error()
			} else {
				msg = "terjadi kesalahan pada sistem server"
			}
			return errors.New(msg)
		}

		return nil
	}
	res, err := guc.qry.GetByID(uint(userID))
	if err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada sistem server"
		}
		return errors.New(msg)
	}

	imageURL, err := helper.UploadTeacherProfilePhotoS3(*avatar, res.Email)
	if err != nil {
		log.Println(err)
		var msg string
		if strings.Contains(err.Error(), "kesalahan input") {
			msg = err.Error()
		} else {
			msg = "gagal upload gambar karena kesalahan pada sistem server"
		}
		return errors.New(msg)
	}
	updateData.Avatar = imageURL

	if err := guc.qry.Update(uint(userID), updateData); err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "tidak ditemukan") {
			msg = err.Error()
		} else {
			msg = "terjadi kesalahan pada sistem server"
		}
		return errors.New(msg)
	}

	return nil
}
