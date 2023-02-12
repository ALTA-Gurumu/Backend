package handler

import (
	"Gurumu/config"
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"context"
	"encoding/json"
	"log"
	"net/http"

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
		client_id := config.GOOGLE_OAUTH_CLIENT_ID1
		project := config.GOOGLE_PROJECT_ID1
		secret := config.GOOGLE_OAUTH_CLIENT_SECRET1

		b := `{"web":{"client_id":"` + client_id + `","project_id":"` + project + `","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"` + secret + `","redirect_uris":["http://localhost:8000/callback"]}}`
		bt := []byte(b)
		config, err := google.ConfigFromJSON(bt, calendar.CalendarEventsScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		code := c.QueryParam("code")
		token, err := config.Exchange(context.Background(), code)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		helper.SaveToken("helper/temporary/token.json", token)

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
