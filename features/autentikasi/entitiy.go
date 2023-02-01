package autentikasi

import "github.com/labstack/echo/v4"

type Core struct {
	ID       uint
	Nama     string
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
}

type AutentikasiHandler interface {
	Login() echo.HandlerFunc
}

type AutentikasiService interface {
	Login(email, password string) (string, Core, error)
}
type AutentikasiData interface {
	Login(email string) (Core, error)
}
