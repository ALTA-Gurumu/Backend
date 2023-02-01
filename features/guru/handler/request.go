package handler

import "Gurumu/features/guru"

type LoginRequest struct {
	Nama     string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type UpdateRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	City     string `json:"city" form:"city"`
	Alamat   string `json:"alamat" form:"alamat"`
	Telepon  string `json:"telepon" form:"telepon"`
}

func ReqToCore(data interface{}) *guru.Core {
	res := guru.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Nama = cnv.Nama
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Nama = cnv.Nama
		res.Password = cnv.Password
		res.Email = cnv.Email
		res.Alamat = cnv.Alamat
		res.Telepon = cnv.Telepon
	default:
		return nil
	}

	return &res
}
