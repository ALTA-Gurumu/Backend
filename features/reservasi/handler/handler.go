package handler

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reservasiHandler struct {
	srv reservasi.ReservasiService
}

func New(srv reservasi.ReservasiService) reservasi.ReservasiHandler {
	return &reservasiHandler{
		srv: srv,
	}
}

func (rh *reservasiHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		data := AddReservasiRequest{}

		if err := c.Bind(&data); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := rh.srv.Add(token, *ToCore(data))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses reservasi guru", ToAddReservasiResponse(res)))
	}
}