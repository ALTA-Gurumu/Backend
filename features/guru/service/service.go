package service

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"
	"math"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
)

type guruUseCase struct {
	qry guru.GuruData
	vld *validator.Validate
}

func New(gd guru.GuruData, vld *validator.Validate) guru.GuruService {
	return &guruUseCase{
		qry: gd,
		vld: vld,
	}
}

// Register implements guru.GuruService
func (guc *guruUseCase) Register(newGuru guru.Core) (guru.Core, error) {
	hashed, err := helper.GeneratePassword(newGuru.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return guru.Core{}, errors.New("password bcrypt process error")
	}

	err = helper.ValidateRegisterRequest(newGuru.Nama, newGuru.Email, newGuru.Password)
	if err != nil {
		msg := "format input tidak lengkap dan/atau kosong"
		return guru.Core{}, errors.New(msg)
	}

	emailCheck := helper.ValidEmail(newGuru.Email)
	if !emailCheck {
		msg := "format email salah"
		return guru.Core{}, errors.New(msg)
	}

	pwdCheck := helper.IsStrongPassword(newGuru.Password)
	if !pwdCheck {
		msg := "password kurang kuat, minimal 6 karakter"
		return guru.Core{}, errors.New(msg)
	}

	namaCheck := helper.IsGoodName(newGuru.Nama)
	if !namaCheck {
		msg := "nama kurang, minimal 6 karakter"
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
			msg = "data not found"
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

	structCheck := helper.IsStructEmpty(updateData)
	if !structCheck {
		return fmt.Errorf("updateData tidak bisa kosong")
	}

	if err := guc.vld.Struct(&updateData); err != nil {
		log.Println(err)
		return fmt.Errorf("validation error: %s", helper.ValidationErrorHandle(err))
	}

	if updateData.Email == "" {
		res, _ := guc.qry.GetByID(uint(userID))

		updateData.Email = res.(guru.Core).Email
	}

	if avatar != nil {
		path, _ := helper.UploadTeacherProfilePhotoS3(*avatar, updateData.Email)
		updateData.Avatar = path
	}

	if ijazah != nil {
		path, _ := helper.UploadTeacherCertificateS3(*ijazah, updateData.Email)
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
func (guc *guruUseCase) ProfileBeranda(loc string, subj string, page int) (map[string]interface{}, []guru.Core, error) {
	if page < 1 {
		page = 1
	}
	limit := 6

	offset := (page - 1) * limit

	totalRecord, res, err := guc.qry.GetBeranda(loc, subj, limit, offset)
	if err != nil {
		log.Println(err)
		return nil, nil, errors.New("internal server error")
	}

	totalPage := int(math.Ceil(float64(totalRecord) / float64(limit)))

	pagination := make(map[string]interface{})
	pagination["page"] = page
	pagination["limit"] = limit
	pagination["offset"] = offset
	pagination["totalRecord"] = totalRecord
	pagination["totalPage"] = totalPage

	return pagination, res, nil
}
