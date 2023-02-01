package handler

import "Gurumu/features/autentikasi"

type AutentikasiResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func ToResponses(data autentikasi.Core) AutentikasiResponse {
	return AutentikasiResponse{
		Nama:  data.Nama,
		Email: data.Email,
		Role:  data.Role,
	}
}
