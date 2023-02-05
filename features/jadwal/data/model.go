package data

import (
	"Gurumu/features/jadwal"

	"gorm.io/gorm"
)

type Guru struct {
	gorm.Model
	Email     string
	Password  string
	Nama      string
	Telepon   string
	Deskripsi string
	Ijazah    string
	Pelajaran string
	Alamat    string
	Avatar    string
	Role      string
	Jadwals   []Jadwal `gorm:"foreignKey:GuruID;references:ID"`
}
type Jadwal struct {
	gorm.Model
	GuruID  uint
	Tanggal string
	Jam     string
	Status  string
}

type JadwalNG struct {
	ID      uint
	Tanggal string
	Jam     string
	Status  string
}

type GuruJadwal struct {
	ID      uint
	Tanggal string
	Jam     string
	Status  string
}

func ToCore(data Jadwal) jadwal.Core {
	return jadwal.Core{
		ID:      data.ID,
		GuruID:  data.GuruID,
		Tanggal: data.Tanggal,
		Jam:     data.Jam,
		Status:  data.Status,
	}
}

func CoreToData(data jadwal.Core) Jadwal {
	return Jadwal{
		Model:   gorm.Model{ID: data.ID},
		GuruID:  data.GuruID,
		Tanggal: data.Tanggal,
		Jam:     data.Jam,
	}
}

func GuruToCore(gj GuruJadwal) jadwal.Core {
	return jadwal.Core{
		ID:      gj.ID,
		Tanggal: gj.Tanggal,
		Jam:     gj.Jam,
		Status:  gj.Status,
	}
}

func ListToCore(data []Jadwal) []jadwal.Core {
	listJadwal := []jadwal.Core{}
	for _, jadwal := range data {
		listJadwal = append(listJadwal, ToCore(jadwal))
	}

	return listJadwal
}
