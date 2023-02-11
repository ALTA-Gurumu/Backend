package handler

import (
	"Gurumu/features/jadwal"
	"Gurumu/helper"
	"log"
	"net/http"
	"strings"

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
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ToResponse(res),
			"message": "berhasil menambahkan jadwal",
		})
	}
}

func (jh *jadwalHandler) GetJadwal() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := jh.srv.GetJadwal(token)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("data tidak ditemukan"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("terdapat masalah pada server"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    GetJadwalResponse(res),
			"message": "sukses menampilkan daftar jadwal",
		})
	}
}
