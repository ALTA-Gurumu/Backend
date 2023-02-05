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
func (guc *guruUseCase) Profile(id uint) (interface{}, error) {
	userID := id
	if userID <= 0 {
		return guru.Core{}, errors.New("token tidak valid")
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
		return guru.Core{}, errors.New(msg)
	}

	return res, nil
}

// Update implements guru.GuruService
func (guc *guruUseCase) Update(token interface{}, updateData guru.Core, avatar *multipart.FileHeader, ijazah *multipart.FileHeader) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return fmt.Errorf("token tidak valid")
	}

	if err := guc.vld.Struct(&updateData); err != nil {
		log.Println(err)
		return fmt.Errorf("validation error: %s", helper.ValidationErrorHandle(err))
	}

	if avatar != nil {
		path, _ := helper.UploadTeacherProfilePhotoS3(*avatar, updateData.Email)
		updateData.Avatar = path
	}

	if ijazah != nil {
		path, _ := helper.UploadTeacherProfilePhotoS3(*ijazah, updateData.Email)
		updateData.Ijazah = path
	}

	if err := guc.qry.Update(uint(userID), updateData); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "tidak ditemukan") {
			return fmt.Errorf("data guru tidak ditemukan")
		}
		return fmt.Errorf("gagal update data guru: %s", err)
	}

	return nil
}

// ProfileBeranda implements guru.GuruService
func (guc *guruUseCase) ProfileBeranda() ([]guru.Core, error) {
	res, err := guc.qry.GetBeranda()
	if err != nil {
		log.Println("no result or server error")
		return []guru.Core{}, errors.New("no result or server error")
	}

	return res, nil
}
