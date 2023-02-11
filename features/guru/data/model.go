package data

import (
	"Gurumu/features/guru"
	"Gurumu/features/jadwal/data"

	"gorm.io/gorm"
)

type Guru struct {
	gorm.Model
	Nama        string
	Email       string
	Password    string
	Telepon     string
	LinkedIn    string
	Gelar       string
	TentangSaya string
	Pengalaman  string
	LokasiAsal  string
	MetodeBljr  string
	Tarif       int
	Pelajaran   string
	Pendidikan  string
	Avatar      string
	Ijazah      string
	Role        string
	Verifikasi  bool
	Latitude    float64       `gorm:"type:decimal(10,8)"`
	Longitude   float64       `gorm:"type:decimal(10,8)"`
	Jadwal      []data.Jadwal `gorm:"foreignKey:GuruID;references:ID"`
}

type GuruRatingBeranda struct {
	ID          uint
	Nama        string
	LokasiAsal  string
	TentangSaya string
	Pelajaran   string
	Avatar      string
	Tarif       int
	Penilaian   float32
}

func RatingToCore(data GuruRatingBeranda) guru.Core {
	return guru.Core{
		ID:          data.ID,
		Nama:        data.Nama,
		TentangSaya: data.TentangSaya,
		LokasiAsal:  data.LokasiAsal,
		Pelajaran:   data.Pelajaran,
		Avatar:      data.Avatar,
		Tarif:       data.Tarif,
		Penilaian:   data.Penilaian,
	}
}

func ListRatingToCore(data []GuruRatingBeranda) []guru.Core {
	listGuru := []guru.Core{}
	for _, v := range data {
		listGuru = append(listGuru, RatingToCore(v))
	}
	return listGuru
}

type Jadwal struct {
	ID      uint
	GuruID  uint
	Tanggal string
	Jam     string
	Status  string
}

func ToCore(data Guru) guru.Core {
	return guru.Core{
		ID:          data.ID,
		Nama:        data.Nama,
		Email:       data.Email,
		Password:    data.Password,
		Telepon:     data.Telepon,
		LinkedIn:    data.LinkedIn,
		Gelar:       data.Gelar,
		TentangSaya: data.TentangSaya,
		Pengalaman:  data.Pengalaman,
		LokasiAsal:  data.LokasiAsal,
		MetodeBljr:  data.MetodeBljr,
		Tarif:       data.Tarif,
		Pelajaran:   data.Pelajaran,
		Pendidikan:  data.Pendidikan,
		Avatar:      data.Avatar,
		Ijazah:      data.Ijazah,
		Role:        data.Role,
		Verifikasi:  data.Verifikasi,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
	}
}

func CoreToData(core guru.Core) Guru {
	return Guru{
		Model:       gorm.Model{ID: core.ID},
		Nama:        core.Nama,
		Email:       core.Email,
		Password:    core.Password,
		Telepon:     core.Telepon,
		LinkedIn:    core.LinkedIn,
		Gelar:       core.Gelar,
		TentangSaya: core.TentangSaya,
		Pengalaman:  core.Pengalaman,
		LokasiAsal:  core.LokasiAsal,
		MetodeBljr:  core.MetodeBljr,
		Tarif:       core.Tarif,
		Pelajaran:   core.Pelajaran,
		Pendidikan:  core.Pendidikan,
		Avatar:      core.Avatar,
		Ijazah:      core.Ijazah,
		Role:        core.Role,
		Verifikasi:  core.Verifikasi,
		Latitude:    core.Latitude,
		Longitude:   core.Longitude,
	}
}
