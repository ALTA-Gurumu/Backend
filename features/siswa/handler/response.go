package handler

import "Gurumu/features/siswa"

type SiswaResponse struct {
	ID       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Nama     string `json:"nama" form:"nama"`
	Telepon  string `json:"telepon" form:"telepon"`
	Alamat   string `json:"alamat" form:"alamat"`
	Avatar   string `json:"avatar" form:"avatar"`
}

type RegisterResponse struct {
	ID    uint   `json:"id" form:"id"`
	Nama  string `json:"nama" form:"nama"`
	Email string `json:"email" form:"email"`
}

func ToResponseRegister(data siswa.Core) RegisterResponse {
	return RegisterResponse{
		ID:    data.ID,
		Nama:  data.Nama,
		Email: data.Email,
	}
}
