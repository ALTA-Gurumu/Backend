package handler

import (
	"Gurumu/features/siswa"
	"Gurumu/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type siswaControl struct {
	srv siswa.SiswaService
}

func New(srv siswa.SiswaService) siswa.SiswaHandler {
	return &siswaControl{
		srv: srv,
	}
}

func (sc *siswaControl) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		// input.Role := "siswa"
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := sc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponseRegister(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil daftar akun baru", dataResp))
	}
}

func (sc *siswaControl) Profile() echo.HandlerFunc {
	return nil
}

func (sc *siswaControl) Delete() echo.HandlerFunc {
	return nil
}

func (sc *siswaControl) Update() echo.HandlerFunc {
	return nil
}
