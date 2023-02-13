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
func TestNotificationTransactionStatus(t *testing.T) {
	data := mocks.NewReservasiData(t)
	// t.Run("success", func(t *testing.T) {

	// 	transactionID := "6ce34ad3-1e3f-463d-b5a2-d420259ce69d"
	// 	srv := New(data)
	// 	data.On("NotificationTransactionStatus", transactionID).Return(nil).Once()
	// 	err := srv.NotificationTransactionStatus(transactionID)
	// 	assert.Nil(t, err)
	// 	data.AssertExpectations(t)
	// })
	t.Run("error check transaction status", func(t *testing.T) {
		srv := New(data)

		err := srv.NotificationTransactionStatus("1234567")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "error check")
		data.AssertExpectations(t)
	})

}

func TestUpdateStatus(t *testing.T) {
	data := mocks.NewReservasiData(t)

	reservasiID := uint(1)
	t.Run("succes update status", func(t *testing.T) {
		srv := New(data)
		data.On("UpdateStatus", uint(1), reservasiID).Return(nil).Once()
		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.UpdateStatus(pToken, uint(1))
		assert.Nil(t, err)
		data.AssertExpectations(t)

	})
	t.Run("not found", func(t *testing.T) {
		srv := New(data)
		data.On("UpdateStatus", uint(1), reservasiID).Return(errors.New("error update reservasi status service")).Once()
		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.UpdateStatus(pToken, uint(1))
		assert.NotNil(t, err)
		data.AssertExpectations(t)

	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		err := srv.UpdateStatus(token, uint(1))
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "valid")

		data.AssertExpectations(t)

	})
}
