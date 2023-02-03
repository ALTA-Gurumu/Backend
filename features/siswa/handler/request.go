package handler

import "Gurumu/features/siswa"

type RegisterRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Nama     string `json:"nama" form:"nama"`
	// Role     string
}

type UpdateRequest struct {
	Nama    string `json:"nama" form:"nama"`
	Email   string `json:"email" form:"email"`
	Alamat  string `json:"alamat" form:"alamat"`
	Telepon string `json:"telepon" form:"telepon"`
	Avatar  string `json:"avatar" form:"avatar"`
}

func ToCore(data interface{}) *siswa.Core {
	res := siswa.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email
		res.Password = cnv.Password
		// res.Role = cnv.Role
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email
		res.Alamat = cnv.Alamat
		res.Telepon = cnv.Telepon
		res.Avatar = cnv.Avatar
	default:
		return nil
	}
	return &res
}
