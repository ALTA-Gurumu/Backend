package jadwal

import "github.com/labstack/echo/v4"

type Core struct {
	ID      uint
	GuruID  uint
	Tanggal string
	Jam     string
	Status  string
}

type JadwalHandler interface {
	Add() echo.HandlerFunc
}
type JadwalService interface {
	Add(token interface{}, newJadwal Core) (Core, error)
}
type JadwalData interface {
	Add(guruID uint, newJadwal Core) (Core, error)
}
