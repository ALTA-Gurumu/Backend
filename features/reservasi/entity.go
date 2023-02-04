package reservasi

import "github.com/labstack/echo/v4"

type Core struct {
	ID               uint
	GuruID           uint
	NamaGuru         string
	SiswaID          uint
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
}

type ReservasiService interface {
	Add(token interface{}, newReservasi Core) (Core, error)
}

type ReservasiData interface {
	Add(siswaID uint, newReservasi Core) (Core, error)
}
