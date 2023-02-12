package service

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
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
