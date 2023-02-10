package service

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	data := mocks.NewReservasiData(t)

	inputData := reservasi.Core{
		GuruID:           1,
		Pesan:            "belajar matematika",
		MetodeBelajar:    "offline",
		Tanggal:          "2023-03-19",
		Jam:              "07.00 WIB",
		AlamatSiswa:      "Jl. Nangka, Mojokerto",
		TeleponSiswa:     "08111",
		MetodePembayaran: "VA Mandiri",
	}

	expectedData := reservasi.Core{
		ID:               1,
		NamaGuru:         "Bejo",
		MetodeBelajar:    "offline",
		Pelajaran:        "Matematika",
		TotalTarif:       75000,
		AlamatSiswa:      "Jl. Nangka, Mojokerto",
		TeleponSiswa:     "08111",
		KodeTransaksi:    "gurumu0101",
		MetodePembayaran: "VA Mandiri",
		NomerVa:          "9898788789675",
		KodeQr:           "kjl432r245",
		BankPenerima:     "Mandiri",
		StatusPembayaran: "sukses",
		TautanGmet:       "gmeet.nkddj/kdj",
		Status:           "ongoing",
	}

	siswaId := uint(2)

	t.Run("success add reservasi", func(t *testing.T) {
		data.On("Add", siswaId, inputData).Return(expectedData, nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Add", siswaId, inputData).Return(reservasi.Core{}, errors.New("not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "tidak ditemukan")
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Add", siswaId, inputData).Return(reservasi.Core{}, errors.New("server problem")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}

func TestMysession(t *testing.T) {
	data := mocks.NewReservasiData(t)

	expectedData := []reservasi.Core{
		{
			ID:         1,
			NamaGuru:   "Bejo",
			Tanggal:    "2023-03-19",
			Jam:        "07.00 PM",
			TautanGmet: "gmeet.jljlaffa",
			Status:     "Selesai",
		}, {
			ID:         2,
			NamaGuru:   "Dono",
			Tanggal:    "2023-03-20",
			Jam:        "06.00 PM",
			TautanGmet: "gmeet.jljlaffa",
			Status:     "Selesai",
		},
	}

	role := "siswa"
	reservasiStatus := "selesai"
	userId := uint(2)

	t.Run("success get mysession", func(t *testing.T) {
		data.On("Mysession", userId, role, reservasiStatus).Return(expectedData, nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Mysession(pToken, role, reservasiStatus)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)

	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Mysession", userId, role, reservasiStatus).Return([]reservasi.Core{}, errors.New("not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Mysession(pToken, role, reservasiStatus)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "tidak ditemukan")
		data.AssertExpectations(t)

	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Mysession", userId, role, reservasiStatus).Return([]reservasi.Core{}, errors.New("server problem")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Mysession(pToken, role, reservasiStatus)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)

	})
}
