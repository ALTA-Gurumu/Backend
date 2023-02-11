package service

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"Gurumu/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	data := mocks.NewGuruData(t)
	v := validator.New()
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

		srv := New(data, v)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.Nama, res.Nama)
		data.AssertExpectations(t)
	})

	t.Run("duplicate", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(guru.Core{}, errors.New("duplicated")).Once()
		srv := New(data, v)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "sudah terdaftar")
		assert.Equal(t, res.Nama, "")
		data.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(guru.Core{}, errors.New("server error")).Once()
		srv := New(data, v)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Nama, "")
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewGuruData(t)
	v := validator.New()
	t.Run("success delete", func(t *testing.T) {
		data.On("Delete", uint(1)).Return(nil).Once()

		srv := New(data, v)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data, v)

		_, token := helper.GenerateJWT(1)

		err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Delete", uint(4)).Return(errors.New("data not found")).Once()

		srv := New(data, v)

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
		srv := New(data, v)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	data := mocks.NewGuruData(t)
	v := validator.New()
	expectedData := guru.Core{
		Nama:        "Putra",
		Email:       "putra123@gmail.com",
		Telepon:     "08111",
		LinkedIn:    "ekacahyaputra",
		Gelar:       "S.T",
		TentangSaya: "Ahli di bidang fisika",
		Pengalaman:  "3 tahun mengajar",
		LokasiAsal:  "Mojokerto",
		MetodeBljr:  "offline",
		Tarif:       75000,
		Pelajaran:   "Fisika",
		Pendidikan:  "Teknik Geofisika",
		Avatar:      "try123ok.s3.aws.amazon.com/avatar.jpg",
		Ijazah:      "try123ok.s3.aws.amazon.com/certificate.jpg",
		Latitude:    -7.1234,
		Longitude:   120.2345,
		Jadwal:      nil,
	}

	guruID := uint(1)
	t.Run("success", func(t *testing.T) {
		data.On("GetByID", guruID).Return(expectedData, nil).Once()

		srv := New(data, v)

		res, err := srv.Profile(guruID)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		data.AssertExpectations(t)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("GetByID", guruID).Return(guru.Core{}, errors.New("data not found")).Once()

		srv := New(data, v)

		res, err := srv.Profile(guruID)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "not found")
		data.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		data.On("GetByID", guruID).Return(guru.Core{}, errors.New("server problem")).Once()

		srv := New(data, v)

		res, err := srv.Profile(guruID)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}

func TestProfilBeranda(t *testing.T) {
	data := mocks.NewGuruData(t)
	v := validator.New()
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

	listRes := expectedData
	listRes[1].ID = uint(2)

	t.Run("success get profile beranda", func(t *testing.T) {
		data.On("GetBeranda", "Mojokerto", "Fisika", 4, 0).Return(2, expectedData, nil).Once()

		srv := New(data, v)

		pagination, res, err := srv.ProfileBeranda("Mojokerto", "Fisika", 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, listRes[0].ID, res[0].ID)
		assert.NotNil(t, pagination)
		data.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		data.On("GetBeranda", "Mojokerto", "Fisika", 4, 0).Return(0, nil, errors.New("not found")).Once()

		srv := New(data, v)

		pagination, res, err := srv.ProfileBeranda("Mojokerto", "Fisika", 1)
		assert.NotNil(t, err)
		assert.Nil(t, pagination)
		assert.Empty(t, res)
		data.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	data := mocks.NewGuruData(t)
	v := validator.New()
	inputData := guru.Core{
		Nama:        "Putra",
		Email:       "putra123@gmail.com",
		Password:    "putra123",
		Telepon:     "08111",
		LinkedIn:    "ekacahyaputra",
		Gelar:       "S.T",
		TentangSaya: "Ahli di bidang Fisika",
		Pengalaman:  "3 tahun mengajar",
		LokasiAsal:  "Mojokerto",
		MetodeBljr:  "offline",
		Tarif:       75000,
		Pelajaran:   "Fisika",
		Pendidikan:  "Teknik Geofisika",
		Avatar:      "try123ok.s3.aws.amazon.com/avatar.jpg",
		Ijazah:      "try123ok.s3.aws.amazon.com/certificate.jpg",
		Latitude:    -7.12334,
		Longitude:   120.9384,
	}

	guruID := uint(1)

	var avatar, ijazah *multipart.FileHeader
	t.Run("success update", func(t *testing.T) {
		data.On("Update", guruID, inputData).Return(nil).Once()

		srv := New(data, v)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Update(pToken, inputData, avatar, ijazah)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("gagal update", func(t *testing.T) {
		data.On("Update", guruID, inputData).Return(errors.New("data tidak ditemukan")).Once()

		srv := New(data, v)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Update(pToken, inputData, avatar, ijazah)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "ditemukan")
		data.AssertExpectations(t)
	})

	t.Run("gagal update", func(t *testing.T) {
		data.On("Update", guruID, inputData).Return(errors.New("masalah pada server")).Once()

		srv := New(data, v)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Update(pToken, inputData, avatar, ijazah)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})

	t.Run("token tidak valid", func(t *testing.T) {
		srv := New(data, v)

		_, token := helper.GenerateJWT(1)

		err := srv.Update(token, inputData, avatar, ijazah)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak valid")
	})
}
