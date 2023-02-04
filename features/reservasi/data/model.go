package data

import (
	"Gurumu/features/reservasi"

	"gorm.io/gorm"
)

type Reservasi struct {
	gorm.Model
	GuruID           uint
	SiswaID          uint
	JadwalID         uint
	Pesan            string
	MetodeBelajar    string
	KodeTransaksi    string
	MetodePembayaran string
	NomerVa          string
	KodeQr           string
	BankPenerima     string
	StatusPembayaran string
	TotalTarif       int
	TautanGmet       string
	Status           string
}

type Guru struct {
	gorm.Model
	Email     string
	Password  string
	Nama      string
	Telepon   string
	Tarif     int
	Deskripsi string
	Ijazah    string
	Pelajaran string
	Alamat    string
	Avatar    string
	Role      string
}
type Jadwal struct {
	gorm.Model
	GuruID  uint
	Tanggal string
	Jam     string
	Status  string
}

func ToCore(data Reservasi) reservasi.Core {
	return reservasi.Core{
		ID:               data.ID,
		GuruID:           data.GuruID,
		SiswaID:          data.SiswaID,
		JadwalID:         data.JadwalID,
		Pesan:            data.Pesan,
		MetodeBelajar:    data.MetodeBelajar,
		KodeTransaksi:    data.KodeTransaksi,
		MetodePembayaran: data.MetodePembayaran,
		NomerVa:          data.NomerVa,
		KodeQr:           data.KodeQr,
		BankPenerima:     data.BankPenerima,
		StatusPembayaran: data.Status,
		TotalTarif:       data.TotalTarif,
		TautanGmet:       data.TautanGmet,
		Status:           data.Status,
	}
}
func CoreToData(data reservasi.Core) Reservasi {
	return Reservasi{
		Model:            gorm.Model{},
		GuruID:           data.GuruID,
		SiswaID:          data.SiswaID,
		JadwalID:         data.JadwalID,
		Pesan:            data.Pesan,
		MetodeBelajar:    data.MetodeBelajar,
		KodeTransaksi:    data.KodeTransaksi,
		MetodePembayaran: data.MetodePembayaran,
		NomerVa:          data.NomerVa,
		KodeQr:           data.KodeQr,
		BankPenerima:     data.BankPenerima,
		StatusPembayaran: data.StatusPembayaran,
		TotalTarif:       data.TotalTarif,
		TautanGmet:       data.TautanGmet,
		Status:           data.Status,
	}
}
