package service

import (
	"Gurumu/features/ulasan"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAdd(t *testing.T) {
	data := mocks.NewUlasanData(t)

	inputData := ulasan.Core{
		Ulasan:    "penjelasannya menarik dan terperinci",
		Penilaian: 5,
	}

	t.Run("success add ulasan", func(t *testing.T) {
		siswaId := uint(1)
		guruId := uint(2)
		data.On("Add", siswaId, guruId, mock.Anything).Return(nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Add(pToken, guruId, inputData)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("data already exist", func(t *testing.T) {
		siswaId := uint(1)
		guruId := uint(2)
		data.On("Add", siswaId, guruId, mock.Anything).Return(errors.New("duplicated")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Add(pToken, guruId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "already exist")
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		siswaId := uint(1)
		guruId := uint(2)
		data.On("Add", siswaId, guruId, mock.Anything).Return(errors.New("server error")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Add(pToken, guruId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		// siswaId := uint(1)
		guruId := uint(2)
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		err := srv.Add(token, guruId, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak valid")
	})
}

func TestGetAll(t *testing.T) {
	data := mocks.NewUlasanData(t)

	expectedData := []ulasan.Core{
		{
			ID:        1,
			NamaGuru:  "Beni",
			Penilaian: 5,
			Ulasan:    "cara menerangkannya sangat jelas",
		}, {
			ID:        2,
			NamaGuru:  "Bejo",
			Penilaian: 4,
			Ulasan:    "gurunya punya banyak cara jitu",
		},
	}

	t.Run("success get all", func(t *testing.T) {
		data.On("GetAll").Return(expectedData, nil).Once()

		srv := New(data)

		_, err := srv.GetAll()
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		data.On("GetAll").Return([]ulasan.Core{}, errors.New("data not found")).Once()

		srv := New(data)

		_, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("GetAll").Return([]ulasan.Core{}, errors.New("server problem")).Once()

		srv := New(data)

		_, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}
