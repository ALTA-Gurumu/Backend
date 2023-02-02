package siswa

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Nama     string `validate:"required"`
	Telepon  string
	Alamat   string
	Avatar   string
	Role     string
}

type SiswaHandler interface {
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type SiswaService interface {
	Register(newStudent Core) (Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updateData Core, avatar *multipart.FileHeader) error
	Delete(token interface{}) error
}

type SiswaData interface {
	Register(newStudent Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) error
	Delete(id uint) error
}
