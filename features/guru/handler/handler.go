package handler

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
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
	return func(c echo.Context) error {
		str := c.Param("guru_id")
		guruID, err := strconv.Atoi(str)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "id guru salah",
			})
		}

		if guruID <= 0 {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "id guru salah",
			})
		}

		res, err := gc.srv.Profile(uint(guruID))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil lihat profil guru", GuruByID(res.(guru.Core))))

	}
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

		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mendaftarkan profil guru", ToResponse(res)))
	}
}

// Update implements guru.GuruHandler
func (gc *guruControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		updateGuru := UpdateRequest{}
		if err := c.Bind(&updateGuru); err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))

		}

		avatar, _ := c.FormFile("avatar")
		ijazah, _ := c.FormFile("ijazah")

		guruCore := guru.Core{}
		copier.Copy(&guruCore, &updateGuru)
		if err := gc.srv.Update(token, guruCore, avatar, ijazah); err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(helper.PrintSuccessReponse(200, "sukses update data guru"))
	}
}

// ProfileBeranda implements guru.GuruHandler
func (gc *guruControl) ProfileBeranda() echo.HandlerFunc {
	return func(c echo.Context) error {
		str := c.QueryParam("page")
		page, _ := strconv.Atoi(str)

		loc := c.QueryParam("lokasi")
		subj := c.QueryParam("pelajaran")

		paginate, res, err := gc.srv.ProfileBeranda(loc, subj, page)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		guruListResp := []ProfileHomeResp{}
		if err := copier.Copy(&guruListResp, &res); err != nil {
			log.Println("guru beranda list:", err)
			return c.JSON(helper.PrintErrorResponse("failed to unmarshall request"))
		}

		pagination := helper.PaginationResponse{
			Page:        paginate["page"].(int),
			Limit:       paginate["limit"].(int),
			Offset:      paginate["offset"].(int),
			TotalRecord: paginate["totalRecord"].(int),
			TotalPage:   paginate["totalPage"].(int),
		}

		response := helper.WithPagination{
			Pagination: pagination,
			Data:       guruListResp,
			Message:    "berhasil tampilkan guru beranda",
		}

		return c.JSON(200, response)
	}
}
