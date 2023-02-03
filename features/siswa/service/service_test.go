package service

import (
	"Gurumu/features/siswa"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/golang-jwt/jwt"
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

func TestProfile(t *testing.T) {
	data := mocks.NewSiswaData(t)

	expectedData := siswa.Core{
		ID:      1,
		Email:   "budi123@gmail.com",
		Nama:    "budi",
		Telepon: "08123456789",
		Alamat:  "Jl. Nangka, Mojokerto, Jawa Timur",
		Avatar:  "https://capstonegurumu.s3.ap-southeast-1.amazonaws.com/files/siswa/1/avatar.jpg",
	}

	t.Run("success show profile", func(t *testing.T) {
		data.On("Profile", uint(1)).Return(expectedData, nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak valid")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Profile", uint(4)).Return(siswa.Core{}, errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		data.On("Profile", mock.Anything).Return(siswa.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewSiswaData(t)
	t.Run("success delete", func(t *testing.T) {
		data.On("Delete", uint(1)).Return(nil).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak valid")
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Delete", uint(4)).Return(errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		data.On("Delete", mock.Anything).Return(errors.New("terdapat masalah pada server")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	data := mocks.NewSiswaData(t)

	inputData := siswa.Core{
		Nama:    "putra",
		Email:   "putra123@gmail.com",
		Alamat:  "Mojokerto",
		Telepon: "08123456789",
		Avatar:  "https://try123ok.s3.ap-southeast-1.amazonaws.com/files/siswa/putra123@gmail.com/avatar.jpeg",
	}

	var avatar *multipart.FileHeader

	t.Run("update success", func(t *testing.T) {
		data.On("Update", uint(1), mock.Anything).Return(nil).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Update(pToken, inputData, avatar)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		err := srv.Update(token, inputData, avatar)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak valid")
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Update", uint(1), inputData).Return(errors.New("data not found")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Update(pToken, inputData, avatar)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

	t.Run("masalah pada server", func(t *testing.T) {
		data.On("Update", uint(1), inputData).Return(errors.New("server problem")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Update(pToken, inputData, avatar)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}
