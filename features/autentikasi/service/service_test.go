package service

import (
	"Gurumu/features/autentikasi"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	data := mocks.NewAutentikasiData(t)

	passwordSiswa, _ := helper.GeneratePassword("han123")

	inputDataSiswa := autentikasi.Core{
		Email:    "han@gmail.com",
		Password: passwordSiswa,
	}
	passwordGuru, _ := helper.GeneratePassword("herdi123")
	inputDataGuru := autentikasi.Core{
		Email:    "herdi@alta.id",
		Password: passwordGuru,
	}

	expectedSiswa := autentikasi.Core{
		ID:       uint(1),
		Nama:     "hannn",
		Email:    "han@gmail.com",
		Password: passwordSiswa,
		Role:     "siswa",
	}

	expectedGuru := autentikasi.Core{
		ID:         uint(1),
		Nama:       "Herdi",
		Email:      "herdi@alta.id",
		Password:   passwordGuru,
		Role:       "guru",
		Verifikasi: false,
	}

	srv := New(data)

	t.Run("sukses login siswa", func(t *testing.T) {
		data.On("Login", inputDataSiswa.Email).Return(expectedSiswa, nil).Once()

		token, res, err := srv.Login(inputDataSiswa.Email, "han123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, expectedSiswa.Nama, res.Nama)
		data.AssertExpectations(t)
	})

	t.Run("sukses login guru", func(t *testing.T) {
		data.On("Login", inputDataGuru.Email).Return(expectedGuru, nil).Once()

		token, res, err := srv.Login(inputDataGuru.Email, "herdi123")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, expectedGuru.Nama, res.Nama)
		data.AssertExpectations(t)
	})

	t.Run("Belum Register", func(t *testing.T) {
		inputEmail := "Jhony@gmail.com"
		data.On("Login", inputEmail).Return(autentikasi.Core{}, errors.New("record not found")).Once()

		token, res, err := srv.Login(inputEmail, "jhonny123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("password salah", func(t *testing.T) {
		data.On("Login", inputDataSiswa.Email).Return(expectedSiswa, nil)

		srv := New(data)
		token, res, err := srv.Login(inputDataSiswa.Email, "be1423")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		data.On("Login", inputDataSiswa.Email).Return(autentikasi.Core{}, errors.New("server error")).Once()
		_, res, err := srv.Login(inputDataSiswa.Email, "hannn123")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}
