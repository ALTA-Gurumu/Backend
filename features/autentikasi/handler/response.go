package handler

import (
	"Gurumu/features/autentikasi"
)

type SiswaAutentikasiResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type GuruAutentikasiResponse struct {
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Verifikasi bool   `json:"verifikasi"`
}

func GuruToResponses(data autentikasi.Core) GuruAutentikasiResponse {
	return GuruAutentikasiResponse{
		Nama:       data.Nama,
		Email:      data.Email,
		Role:       data.Role,
		Verifikasi: data.Verifikasi,
	}

}

func SiswaToResponses(data autentikasi.Core) SiswaAutentikasiResponse {
	return SiswaAutentikasiResponse{
		Nama:  data.Nama,
		Email: data.Email,
		Role:  data.Role,
	}

}
