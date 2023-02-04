package handler

import "Gurumu/features/reservasi"

type AddReservasiResponse struct {
	ID            uint   `json:"id" form:"id"`
	NamaGuru      string `json:"nama_guru" form:"nama_guru"`
	MetodeBelajar string `json:"metode_belajar" form:"metode_belajar"`
	Pelajaran     string `json:"pelajaran" form:"pelajaran"`
	TotalTarif    int    `json:"total_tarif" form:"total_tarif"`
	AlamatSiswa   string `json:"alamat_siswa" form:"alamat_siswa"`
	TeleponSiswa  string `json:"telepon_siswa" form:"telepon_siswa"`
}

func ToAddReservasiResponse(data reservasi.Core) AddReservasiResponse {
	return AddReservasiResponse{
		ID:            data.ID,
		NamaGuru:      data.NamaGuru,
		MetodeBelajar: data.MetodeBelajar,
		Pelajaran:     data.Pelajaran,
		TotalTarif:    data.TotalTarif,
		AlamatSiswa:   data.AlamatSiswa,
		TeleponSiswa:  data.TeleponSiswa,
	}
}
