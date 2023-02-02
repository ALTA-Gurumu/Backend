package handler

import (
	"Gurumu/features/jadwal"
	"Gurumu/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jadwalHandler struct {
	srv jadwal.JadwalService
}

func New(srv jadwal.JadwalService) jadwal.JadwalHandler {
	return &jadwalHandler{
		srv: srv,
	}
}

func (jh *jadwalHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputData := JadwalRequest{}
		if err := c.Bind(&inputData); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := jh.srv.Add(c.Get("user"), *ToCore(inputData))
		if err != nil {
			log.Println("error running add jadwal service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ToResponse(res),
			"message": "berhasil menambahkan jadwal",
		})
	}
}
