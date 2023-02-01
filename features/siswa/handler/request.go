package handler

import "Gurumu/features/siswa"

type RegisterRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Nama     string `json:"nama" form:"nama"`
	// Role     string
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
	default:
		return nil
	}
	return &res
}
