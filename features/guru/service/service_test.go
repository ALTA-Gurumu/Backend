package service

import (
	"Gurumu/features/guru"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewGuruData(t)

	inputData := guru.Core{
		Nama:     "Putra",
		Email:    "putra123@gmail.com",
		Password: "putra123",
	}

	expectedData := guru.Core{
		Nama:  "Putra",
		Email: "putra123@gmail.com",
	}

	t.Run("success register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(expectedData, nil).Once()

		srv := New(data)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.Nama, res.Nama)
		data.AssertExpectations(t)
	})

	t.Run("duplicate", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(guru.Core{}, errors.New("duplicated")).Once()
		srv := New(data)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "sudah terdaftar")
		assert.Equal(t, res.Nama, "")
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(guru.Core{}, errors.New("server error")).Once()
		srv := New(data)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Nama, "")
		data.AssertExpectations(t)
	})
}
