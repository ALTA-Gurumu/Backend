package guru

import (
	"Gurumu/features/jadwal/data"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Nama        string
	Email       string
	Password    string
	Telepon     string
	LinkedIn    string
	Gelar       string
	TentangSaya string
	Pengalaman  string
	LokasiAsal  string
	Offline     bool
	Online      bool
	Tarif       string
	Pelajaran   string
	Pendidikan  string
	Avatar      string
	Ijazah      string
	Role        string
	Latitude    string
	Longitude   string
	Jadwal      []data.JadwalNG
	Penilaian   float32
}

type GuruHandler interface {
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	ProfileBeranda() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type GuruService interface {
	Register(newGuru Core) (Core, error)
	Profile(id uint) (interface{}, error)
	ProfileBeranda() ([]Core, error)
	Update(token interface{}, updateData Core, avatar *multipart.FileHeader, ijazah *multipart.FileHeader) error
	Delete(token interface{}) error
}

type GuruData interface {
	Register(newGuru Core) (Core, error)
	GetByID(id uint) (interface{}, error)
	GetBeranda() ([]Core, error)
	Update(id uint, updateData Core) error
	Delete(id uint) error
}
