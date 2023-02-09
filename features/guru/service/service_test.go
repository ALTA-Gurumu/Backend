package service

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"testing"

	"github.com/golang-jwt/jwt"
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

func TestDelete(t *testing.T) {
	data := mocks.NewGuruData(t)

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
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Delete", uint(4)).Return(errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
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

// func TestProfile(t *testing.T) {
// 	data := mocks.NewGuruData(t)

// 	expectedData := guru.Core{
// 		Nama: "Putra",
// 		Email: "putra123@gmail.com",
// 		Telepon: "08111",
// 		LinkedIn: "ekacahyaputra",
// 		Gelar: "S.T",
// 		TentangSaya: "Ahli di bidang fisika",
// 		Pengalaman: "3 tahun mengajar",
// 		LokasiAsal: "Mojokerto",
// 		MetodeBljr: "offline",
// 		Tarif: 75000,
// 		Pelajaran: "Fisika",
// 		Pendidikan: "Teknik Geofisika",
// 		Avatar: "try123ok.s3.aws.amazon.com/avatar.jpg",
// 		Ijazah: "try123ok.s3.aws.amazon.com/certificate.jpg",
// 		Latitude: "-7.1234",
// 		Longitude: "120.2345",
// 		Jadwal: ,
// 	}
// }

func TestProfilBeranda(t *testing.T) {
	data := mocks.NewGuruData(t)

	expectedData := []guru.Core{
		{
			ID:          1,
			Nama:        "Putra",
			LokasiAsal:  "Mojokerto",
			TentangSaya: "berpengalaman mengajar 3 tahun",
			Pelajaran:   "Fisika",
			Avatar:      "try123ok.s3.aws.amazon.com/avatar.jpg",
			Tarif:       75000,
			Penilaian:   5,
		}, {
			ID:          2,
			Nama:        "Bejo",
			LokasiAsal:  "Mojokerto",
			TentangSaya: "berpengalaman mengajar 5 tahun",
			Pelajaran:   "Fisika",
			Avatar:      "try123ok.s3.aws.amazon.com/avatar.jpg",
			Tarif:       100000,
			Penilaian:   5,
		},
	}

	t.Run("success get profile beranda", func(t *testing.T) {
		data.On("GetBeranda", "Mojokerto", "Fisika").Return(expectedData, nil).Once()

		srv := New(data)

		res, err := srv.ProfileBeranda("Mojokerto", "Fisika")
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		data.On("GetBeranda", "Mojokerto", "Fisika").Return([]guru.Core{}, errors.New("not found")).Once()

		srv := New(data)

		res, err := srv.ProfileBeranda("Mojokerto", "Fisika")
		assert.NotNil(t, err)
		assert.Empty(t, res)
		data.AssertExpectations(t)
	})
}
