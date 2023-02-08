package service

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"google.golang.org/api/calendar/v3"
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

	//contoh panggil fungsi
	helper.CreateEvent(
		&calendar.Event{
			Summary:     "Test Event",
			Location:    "Somewhere",
			Description: "This is a test event.",
			Start: &calendar.EventDateTime{
				DateTime: time.Now().Format(time.RFC3339),
				TimeZone: "Asia/Jakarta",
			},
			End: &calendar.EventDateTime{
				DateTime: time.Now().Add(time.Hour * 2).Format(time.RFC3339),
				TimeZone: "Asia/Jakarta",
			},
		})

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
