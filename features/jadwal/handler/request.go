package handler

import "Gurumu/features/jadwal"

type JadwalRequest struct {
	Tanggal string `json:"tanggal" form:"tanggal"`
	Jam     string `json:"jam" form:"jam"`
}

func ToCore(data JadwalRequest) *jadwal.Core {
	return &jadwal.Core{
		Tanggal: data.Tanggal,
		Jam:     data.Jam,
	}
}
