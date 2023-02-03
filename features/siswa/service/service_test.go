package service

import (
	"Gurumu/features/siswa"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewSiswaData(t)

	inputData := siswa.Core{
		Email:    "budi123@gmail.com",
		Password: "budi123",
		Nama:     "budi",
	}
	expectedData := siswa.Core{
		ID:    1,
		Nama:  "budi",
		Email: "budi123@gmail.com",
	}

	t.Run("success register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(expectedData, nil).Once()

		srv := New(data)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		assert.Equal(t, expectedData.Nama, res.Nama)
		data.AssertExpectations(t)
	})

	t.Run("duplicate", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(siswa.Core{}, errors.New("duplicated")).Once()
		srv := New(data)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "already exist")
		assert.Equal(t, res.Nama, "")
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(siswa.Core{}, errors.New("server error")).Once()
		srv := New(data)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Nama, "")
		data.AssertExpectations(t)
	})
}
