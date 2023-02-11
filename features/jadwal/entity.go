package jadwal

import "github.com/labstack/echo/v4"

type Core struct {
	ID      uint
	GuruID  uint
	Tanggal string `validate:"required"`
	Jam     string `validate:"required"`
	Status  string
}

type JadwalHandler interface {
	Add() echo.HandlerFunc
	GetJadwal() echo.HandlerFunc
}
type JadwalService interface {
	Add(token interface{}, newJadwal Core) (Core, error)
	GetJadwal(token interface{}) ([]Core, error)
}
type JadwalData interface {
	Add(guruID uint, newJadwal Core) (Core, error)
	GetJadwal(guruID uint) ([]Core, error)
}
