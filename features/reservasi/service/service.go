package service

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"
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

	err := helper.ValidatitonReservasiRequest(newReservasi.MetodeBelajar, newReservasi.Tanggal, newReservasi.Jam, newReservasi.MetodePembayaran)
	if err != nil {
		return reservasi.Core{}, errors.New("kesalahan input dari sisi user")
	}
	res, err := rs.qry.Add(uint(siswaID), newReservasi)
	if err != nil {
		return reservasi.Core{}, err
	}
	return res, nil

}

func (rs *reservasiService) Mysession(token interface{}, role, reservasiStatus string) ([]reservasi.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return []reservasi.Core{}, fmt.Errorf("token tidak valid")
	}

	if role != "guru" && role != "siswa" {
		return []reservasi.Core{}, fmt.Errorf("role/status tidak valid")
	}

	if reservasiStatus != "selesai" && reservasiStatus != "ongoing" && reservasiStatus != "" {
		return []reservasi.Core{}, fmt.Errorf("role/status tidak valid")
	}

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

func (rs *reservasiService) UpdateStatus(token interface{}, reservasiID uint) error {

	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return fmt.Errorf("token tidak valid")
	}

	err := rs.qry.UpdateStatus(uint(userID), reservasiID)
	if err != nil {
		log.Println("error update reservasi status in service: ", err.Error())
		return errors.New("error update reservasi status service")
	}
	return nil
}
