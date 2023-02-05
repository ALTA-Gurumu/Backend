package handler

import "Gurumu/features/jadwal"

type JadwalResponse struct {
	ID      uint   `json:"id" form:"id"`
	GuruID  uint   `json:"guru_id" form:"guru_id"`
	Tanggal string `json:"tanggal" form:"tanggal"`
	Jam     string `json:"jam" form:"jam"`
}

func ToResponse(data jadwal.Core) JadwalResponse {
	return JadwalResponse{
		ID:      data.ID,
		GuruID:  data.GuruID,
		Tanggal: data.Tanggal,
		Jam:     data.Jam,
	}
}

func GetJadwalResponse(data []jadwal.Core) []JadwalResponse {
	listResp := []JadwalResponse{}
	for _, jadwal := range data {
		listResp = append(listResp, ToResponse(jadwal))
	}
	return listResp
}
