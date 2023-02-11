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
func (rs *reservasiService) CheckPaymentStatus(kodeTransaksi string) (string, error) {
	res, err := helper.CheckStatusPayment(kodeTransaksi)
	paymentStatus := ""
	if err != nil {

		return paymentStatus, err
	}
	if res.TransactionStatus == "settlement" {
		paymentStatus = "Sukses"

	}

	return paymentStatus, nil
}
func (rs *reservasiService) Add(token interface{}, newReservasi reservasi.Core) (reservasi.Core, error) {
	siswaID := helper.ExtractToken(token)
	res, err := rs.qry.Add(uint(siswaID), newReservasi, rs.CheckPaymentStatus)

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
			Summary:     "Konsultasi gurumu",
			Location:    "",
			Description: "matematika.",
			ConferenceData: &calendar.ConferenceData{
				CreateRequest: &calendar.CreateConferenceRequest{
					RequestId: "sfsfs",
					ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
						Type: "hangoutsMeet"},
					Status: &calendar.ConferenceRequestStatus{
						StatusCode: "success"},
				}},

			Start: &calendar.EventDateTime{
				DateTime: time.Now().Add(time.Hour * 1).Format(time.RFC3339),
				TimeZone: "Asia/Jakarta",
			},
			End: &calendar.EventDateTime{
				DateTime: time.Now().Add(time.Hour * 4).Format(time.RFC3339),
				TimeZone: "Asia/Jakarta",
			},
			Attendees: []*calendar.EventAttendee{
				{Email: "herdiladania11@gmail.com"},
				{Email: "sucinascaisar@gmail.com"},
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
