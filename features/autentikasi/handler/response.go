package handler

import (
	"Gurumu/features/autentikasi"
)

type SiswaAutentikasiResponse struct {
	ID    uint   `json:"siswa_id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type GuruAutentikasiResponse struct {
	ID         uint   `json:"guru_id"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Verifikasi bool   `json:"verifikasi"`
}

func GuruToResponses(data autentikasi.Core) GuruAutentikasiResponse {
	return GuruAutentikasiResponse{
		ID:         data.ID,
		Nama:       data.Nama,
		Email:      data.Email,
		Role:       data.Role,
		Verifikasi: data.Verifikasi,
	}

}

func SiswaToResponses(data autentikasi.Core) SiswaAutentikasiResponse {
	return SiswaAutentikasiResponse{
		ID:    data.ID,
		Nama:  data.Nama,
		Email: data.Email,
		Role:  data.Role,
	}

}
