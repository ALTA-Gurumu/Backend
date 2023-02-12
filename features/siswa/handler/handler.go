package handler

import (
	"Gurumu/features/siswa"
	"Gurumu/helper"
	"log"
	"mime/multipart"
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
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "format inputan salah",
			})
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
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := sc.srv.Profile(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponseProfil(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", dataResp))
	}
}

func (sc *siswaControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		var avatar *multipart.FileHeader

		updateData := UpdateRequest{}
		err := c.Bind(&updateData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		file, err := c.FormFile("avatar")
		if file != nil && err == nil {
			avatar = file
		} else if file != nil && err != nil {
			log.Println("error baca avatar")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("inputan salah"))
		}

		err2 := sc.srv.Update(token, *ToCore(updateData), avatar)
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil mengganti profil siswa"))
	}
}

func (sc *siswaControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		err := sc.srv.Delete(token)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil hapus akun"))
	}
}
