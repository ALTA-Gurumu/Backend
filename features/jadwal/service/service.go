package service

import (
	"Gurumu/features/jadwal"
	"Gurumu/helper"
	"errors"
	"log"
	"strings"

	"github.com/go-playground/validator"
)

type jadwalService struct {
	qry jadwal.JadwalData
	vld *validator.Validate
}

func New(jd jadwal.JadwalData) jadwal.JadwalService {
	return &jadwalService{
		qry: jd,
		vld: validator.New(),
	}
}
func (js *jadwalService) Add(token interface{}, newJadwal jadwal.Core) (jadwal.Core, error) {
	GuruID := helper.ExtractToken(token)
	if GuruID <= 0 {
		log.Println("pengguna tidak ditemukan")
		return jadwal.Core{}, errors.New("pengguna tidak ditemukan")
	}

	res, err := js.qry.Add(uint(GuruID), newJadwal)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terjadi kesalahan pada server"
		}
		return jadwal.Core{}, errors.New(msg)
	}

	return res, nil
}

func (js *jadwalService) GetJadwal(token interface{}) ([]jadwal.Core, error) {
	guruID := helper.ExtractToken(token)
	if guruID <= 0 {
		log.Println("error extract token")
		return []jadwal.Core{}, errors.New("data tidak ditemukan")
	}

	res, err := js.qry.GetJadwal(uint(guruID))

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		log.Println("error jadwal history service: ", err.Error())
		return []jadwal.Core{}, errors.New(msg)
	}
	return res, nil
}
