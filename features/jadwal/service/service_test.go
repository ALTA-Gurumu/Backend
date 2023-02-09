package service

import (
	"Gurumu/features/jadwal"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	data := mocks.NewJadwalData(t)

	inputData := jadwal.Core{
		Tanggal: "19-03-2023",
		Jam:     "09.00 PM",
	}

	expectedData := jadwal.Core{
		ID:      1,
		GuruID:  1,
		Tanggal: "19-03-2023",
		Jam:     "09.00 PM",
	}

	guruId := uint(1)
	t.Run("success add jadwal", func(t *testing.T) {
		data.On("Add", guruId, inputData).Return(expectedData, nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, res.Tanggal, expectedData.Tanggal)
		data.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Add", guruId, inputData).Return(jadwal.Core{}, errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "tidak ditemukan")
		data.AssertExpectations(t)
	})

	t.Run("terjadi kesalahan pada server", func(t *testing.T) {
		data.On("Add", guruId, inputData).Return(jadwal.Core{}, errors.New("server problem")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		res, err := srv.Add(token, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "pengguna tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})
}

func TestGetJadwal(t *testing.T) {
	data := mocks.NewJadwalData(t)

	expectedData := []jadwal.Core{
		{
			ID:      1,
			GuruID:  1,
			Tanggal: "19-03-2023",
			Jam:     "09.00 PM",
		}, {
			ID:      2,
			GuruID:  1,
			Tanggal: "19-03-2023",
			Jam:     "07.00 PM",
		},
	}

	guruId := uint(1)

	t.Run("success get jadwal", func(t *testing.T) {
		data.On("GetJadwal", guruId).Return(expectedData, nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetJadwal(pToken)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("GetJadwal", guruId).Return([]jadwal.Core{}, errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetJadwal(pToken)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "tidak ditemukan")
		data.AssertExpectations(t)
	})

	t.Run("terdapat masalah pada server", func(t *testing.T) {
		data.On("GetJadwal", guruId).Return([]jadwal.Core{}, errors.New("server problem")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.GetJadwal(pToken)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		res, err := srv.GetJadwal(token)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "tidak ditemukan")
	})
}
