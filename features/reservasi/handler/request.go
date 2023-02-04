package handler

import (
	"Gurumu/features/reservasi"
)

type AddReservasiRequest struct {
	GuruID        uint   `json:"guru_id" form:"guru_id"`
	Pesan         string `json:"pesan" form:"pesan"`
	MetodeBelajar string `json:"metode_belajar" form:"metode_belajar"`
	Tanggal       string `json:"tanggal" form:"tanggal"`
	Jam           string `json:"jam" form:"jam"`
	AlamatSiswa   string `json:"alamat_siswa" form:"alamat_siswa"`
	TeleponSiswa  string `json:"telepon_siswa" form:"telepon_siswa"`
}

func ToCore(data AddReservasiRequest) *reservasi.Core {
	return &reservasi.Core{
		GuruID:        data.GuruID,
		Pesan:         data.Pesan,
		MetodeBelajar: data.MetodeBelajar,
		Tanggal:       data.Tanggal,
		Jam:           data.Jam,
		AlamatSiswa:   data.AlamatSiswa,
		TeleponSiswa:  data.TeleponSiswa,
	}
}
