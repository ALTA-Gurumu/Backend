package service

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

type reservasiService struct {
	qry reservasi.ReservasiData
	vld *validator.Validate
}

func New(rd reservasi.ReservasiData) reservasi.ReservasiService {
	return &reservasiService{
		qry: rd,
		vld: validator.New(),
	}
}

func (rs *reservasiService) Add(token interface{}, newReservasi reservasi.Core) (reservasi.Core, error) {
	siswaID := helper.ExtractToken(token)
	res, err := rs.qry.Add(uint(siswaID), newReservasi)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "internal server error"
		}
		return reservasi.Core{}, errors.New(msg)
	}
	return res, nil

}
