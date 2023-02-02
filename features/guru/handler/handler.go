package handler

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type guruControl struct {
	srv guru.GuruService
}

func New(srv guru.GuruService) guru.GuruHandler {
	return &guruControl{
		srv: srv,
	}
}

// Delete implements guru.GuruHandler
func (gc *guruControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := gc.srv.Delete(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil delete profil guru", err))
	}
}

// Profile implements guru.GuruHandler
func (gc *guruControl) Profile() echo.HandlerFunc {
	panic("unimplemented")
}

// Register implements guru.GuruHandler
func (gc *guruControl) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		res, err := gc.srv.Register(*ReqToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", ToResponse(res)))
	}
}

// Update implements guru.GuruHandler
func (gc *guruControl) Update() echo.HandlerFunc {
	panic("unimplemented")
}
