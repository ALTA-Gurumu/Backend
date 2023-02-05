package handler

import (
	"Gurumu/features/ulasan"
	"Gurumu/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ulasanControl struct {
	srv ulasan.UlasanService
}

func New(srv ulasan.UlasanService) ulasan.UlasanHandler {
	return &ulasanControl{
		srv: srv,
	}
}

func (uc *ulasanControl) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		guruID := c.Param("guruid")
		cnv, err := strconv.Atoi(guruID)
		if err != nil {
			log.Println("add ulasan param error")
			return c.JSON(http.StatusBadRequest, "ID guru salah")
		}

		input := UlasanRegisterRequest{}

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		err2 := uc.srv.Add(token, uint(cnv), *ToCore(input))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil menambahkan ulasan"))
	}
}
func (uc *ulasanControl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.srv.GetAll()
		if err != nil {
			log.Println("handler error")
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		allUlasan := ListAllUlasanToResponse(res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses menampilkan ulasan", allUlasan))
	}
}
func (uc *ulasanControl) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		guruId := c.Param("guruid")
		cnv, err := strconv.Atoi(guruId)
		if err != nil {
			log.Println("get ulasan by id param error")
			return c.JSON(http.StatusBadRequest, "ID guru salah")
		}

		res, err := uc.srv.GetById(uint(cnv))
		if err != nil {
			log.Println("handler error")
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		ulasanGuru := ListUlasanGuruToResponse(res)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses mendapatkan ulasan dan penilaian", ulasanGuru))
	}
}