package handler

import (
	"Gurumu/features/autentikasi"
	"Gurumu/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type autentikasiControll struct {
	srv autentikasi.AutentikasiService
}

func New(srv autentikasi.AutentikasiService) autentikasi.AutentikasiHandler {
	return &autentikasiControll{
		srv: srv,
	}
}

func (ac *autentikasiControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "kesalahan input dari sisi user")
		}

		token, res, err := ac.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponses(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "login sukses", dataResp, token))
	}

}
