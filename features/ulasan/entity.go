package ulasan

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	GuruId    uint
	SiswaId   uint
	Ulasan    string  `validate:"required"`
	Penilaian float32 `validate:"required"`
	NamaSiswa string
	NamaGuru  string
}

type UlasanHandler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetById() echo.HandlerFunc
}

type UlasanService interface {
	Add(token interface{}, guruId uint, newUlasan Core) error
	GetAll() ([]Core, error)
	GetById(guruId uint) ([]Core, error)
}

type UlasanData interface {
	Add(siswaId, guruId uint, newUlasan Core) error
	GetAll() ([]Core, error)
	GetById(guruId uint) ([]Core, error)
}
