package data

import (
	"Gurumu/features/guru"
	"Gurumu/features/jadwal"

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
	Offline     bool
	Online      bool
	Tarif       string
	Pelajaran   string
	Pendidikan  string
	Avatar      string
	Ijazah      string
	Role        string
	Latitude    string
	Longitude   string
	Jadwal      []Jadwal `gorm:"foreignKey:GuruID;references:ID"`
}

type Jadwal struct {
	gorm.Model
	GuruID  uint
	Tanggal string
	Jam     string
	Status  string
}

func JadwalToCore(data Jadwal) jadwal.Core {
	return jadwal.Core{
		ID:      data.ID,
		GuruID:  data.GuruID,
		Tanggal: data.Tanggal,
		Jam:     data.Jam,
		Status:  data.Status,
	}
}

func CoreToJadwal(data jadwal.Core) Jadwal {
	return Jadwal{
		Model:   gorm.Model{ID: data.ID},
		GuruID:  data.GuruID,
		Tanggal: data.Tanggal,
		Jam:     data.Jam,
	}
}

func ListToCore(data []Jadwal) []jadwal.Core {
	listJadwal := []jadwal.Core{}
	for _, jadwal := range data {
		listJadwal = append(listJadwal, JadwalToCore(jadwal))
	}

	return listJadwal
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
		Offline:     false,
		Online:      false,
		Tarif:       data.Tarif,
		Pelajaran:   data.Pelajaran,
		Pendidikan:  data.Pendidikan,
		Avatar:      data.Avatar,
		Ijazah:      data.Ijazah,
		Role:        data.Role,
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
		Offline:     false,
		Online:      false,
		Tarif:       core.Tarif,
		Pelajaran:   core.Pelajaran,
		Pendidikan:  core.Pendidikan,
		Avatar:      core.Avatar,
		Ijazah:      core.Ijazah,
		Role:        core.Role,
		Latitude:    core.Latitude,
		Longitude:   core.Longitude,
	}
}
