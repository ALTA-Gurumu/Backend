package handler

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
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

func (rh *reservasiHandler) Callback() echo.HandlerFunc {
	return func(c echo.Context) error {
		b, _ := os.ReadFile("credentials.json")
		config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		code := c.QueryParam("code")
		// state := c.Param("state")
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		helper.SaveToken("features/reservasi/credentials/token.json", token)

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses kembalikan token"))
	}
}

func (rh *reservasiHandler) Mysession() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		role := c.QueryParam("role")
		reservasiStatus := c.QueryParam("status")

		res, err := rh.srv.Mysession(token, role, reservasiStatus)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		if role == "siswa" {
			return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menampilkan sesi siswa", ToListSesikuSiswaResponse(res)))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "sukses menampilkan sesi guru", ToListSesikuGuruResponse(res)))
	}
}

// CallbackMid implements reservasi.ReservasiHandler
func (rh *reservasiHandler) CallbackMid() echo.HandlerFunc {
	return func(c echo.Context) error {
		str := c.Param("kode")
		err := rh.srv.CallbackMid(str)

		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "sukses deliver status bayar"))
	}
}

func (rh *reservasiHandler) NotificationTransactionStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notificationPayload map[string]interface{}
		err := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		kodeTransaksi, exists := notificationPayload["transaction_id"].(string)
		if !exists {
			return c.JSON(http.StatusBadRequest, err)
		}

		err = rh.srv.NotificationTransactionStatus(kodeTransaksi)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(c.Response().Write([]byte("notifikasi midtrans sukses")))
	}
}
