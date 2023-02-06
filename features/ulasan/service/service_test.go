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
