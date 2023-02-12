package reservasi

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID               uint
	Role             string
	GuruID           uint
	NamaGuru         string
	SiswaID          uint
	NamaSiswa        string
	AlamatSiswa      string
	TeleponSiswa     string
	JadwalID         uint
	Tanggal          string
	Jam              string
	Pesan            string
	Pelajaran        string
	MetodeBelajar    string
	KodeTransaksi    string
	MetodePembayaran string
	NomerVa          string
	KodeQr           string
	BankPenerima     string
	StatusPembayaran string
	TotalTarif       int
	TautanGmet       string
	Status           string
}

type ReservasiHandler interface {
	Add() echo.HandlerFunc
	Mysession() echo.HandlerFunc
	NotificationTransactionStatus() echo.HandlerFunc
	Callback() echo.HandlerFunc
	UpdateStatus() echo.HandlerFunc
}

type ReservasiService interface {
	Add(token interface{}, newReservasi Core) (Core, error)
	Mysession(token interface{}, role, reservasiStatus string) ([]Core, error)
	NotificationTransactionStatus(kodeTransaksi string) error
	UpdateStatus(token interface{}, reservasiID uint) error
}

type ReservasiData interface {
	Add(siswaID uint, newReservasi Core) (Core, error)
	Mysession(userID uint, role, reservasiStatus string) ([]Core, error)
	UpdateDataByTrfID(kode string, updateRes Core) error
	NotificationTransactionStatus(kodeTransaksi, transStatus string) error
	UpdateStatus(userID, reservasiID uint) error
}
