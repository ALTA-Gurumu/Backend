package guru

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Nama      string `validate:"required"`
	Telepon   string
	Deskripsi string
	Ijazah    string
	Pelajaran string
	Alamat    string
	Avatar    string
	Role      string
}

type GuruHandler interface {
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type GuruService interface {
	Register(newGuru Core) (Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updateData Core, avatar *multipart.FileHeader) error
	Delete(token interface{}) error
}

type GuruData interface {
	Register(newGuru Core) (Core, error)
	Profile(id uint) (Core, error)
	Update(id uint, updateData Core) error
	Delete(id uint) error
}
