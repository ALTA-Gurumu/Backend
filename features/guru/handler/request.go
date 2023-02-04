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
	Nama      string `json:"nama" form:"nama"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Telepon   string `json:"telepon" form:"telepon"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Pelajaran string `json:"pelajaran" form:"pelajaran"`
	Alamat    string `json:"alamat" form:"alamat"`
	Avatar    string `json:"avatar" form:"avatar"`
	Ijazah    string `json:"ijazah" form:"ijazah"`
	Role      string
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
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
		res.Email = cnv.Email
		res.Nama = cnv.Nama
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Nama = cnv.Nama
		res.Email = cnv.Email
		res.Password = cnv.Password
		res.Telepon = cnv.Telepon
		res.Deskripsi = cnv.Deskripsi
		res.Pelajaran = cnv.Pelajaran
		res.Alamat = cnv.Alamat
		res.Avatar = cnv.Avatar
		res.Ijazah = cnv.Ijazah
		res.Latitude = cnv.Latitude
		res.Longitude = cnv.Longitude
	default:
		return nil
	}

	return &res
}
