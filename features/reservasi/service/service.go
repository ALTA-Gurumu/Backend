package service

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
	"log"
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

	// //contoh panggil fungsi
	// helper.CreateEvent(
	// 	&calendar.Event{
	// 		Summary:     "Test Event",
	// 		Location:    "Somewhere",
	// 		Description: "This is a test event.",
	// 		Start: &calendar.EventDateTime{
	// 			DateTime: time.Now().Format(time.RFC3339),
	// 			TimeZone: "Asia/Jakarta",
	// 		},
	// 		End: &calendar.EventDateTime{
	// 			DateTime: time.Now().Add(time.Hour * 2).Format(time.RFC3339),
	// 			TimeZone: "Asia/Jakarta",
	// 		},
	// 	})

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

func (rs *reservasiService) Mysession(token interface{}, role, reservasiStatus string) ([]reservasi.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := rs.qry.Mysession(uint(userID), role, reservasiStatus)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "internal server error"
		}
		return []reservasi.Core{}, errors.New(msg)
	}
	return res, nil

}

// CallbackMid implements reservasi.ReservasiService
func (rs *reservasiService) CallbackMid(kode string) error {
	statusBayar := "Terbayar"

	err := rs.qry.UpdateDataByTrfID(kode, reservasi.Core{

		StatusPembayaran: statusBayar,
	})

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "internal server error"
		}
		return errors.New(msg)
	}

	helper.CreateEvent(
		&calendar.Event{
			Summary:     "Test Event",
			Location:    "Somewhere",
			Description: "This is a test event.",
			Start: &calendar.EventDateTime{
				DateTime: time.Now().Add(time.Hour * 2).Format(time.RFC3339),
				TimeZone: "Asia/Jakarta",
			},
			End: &calendar.EventDateTime{
				DateTime: time.Now().Add(time.Hour * 4).Format(time.RFC3339),
				TimeZone: "Asia/Jakarta",
			},
			Attendees: []*calendar.EventAttendee{
				&calendar.EventAttendee{Email: "ariadi.ahmadd@gmail.com"},
			},
			Reminders: &calendar.EventReminders{
				UseDefault: true,
				// Overrides: []*calendar.EventReminder{
				// 	{Method: "email", Minutes: 10},
				// },
			},
		})

	return nil
}
func (rs *reservasiService) NotificationTransactionStatus(kodeTransaksi string) error {

	paymentStatus, err := helper.CheckStatusPayment(kodeTransaksi)
	if err != nil {
		log.Println("error check transaction status: ", err.Error())
		return errors.New("error check transaction status")
	}

	err = rs.qry.NotificationTransactionStatus(kodeTransaksi, paymentStatus.TransactionStatus)
	if err != nil {
		log.Println("error get notificationtransactionstatus data in service: ", err.Error())
		return errors.New("error get notificationtransactionstatus data in service")
	}
	return nil
}
