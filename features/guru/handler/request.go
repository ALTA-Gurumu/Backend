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
	Latitude  string `json:"latitude" form:"latitude"`
	Longitude string `json:"longitude" form:"longitude"`
	Role      string
}

func ReqToCore(data interface{}) *guru.Core {
	res := guru.Core{}

	switch v := data.(type) {
	case LoginRequest:
		res.Nama = v.Nama
		res.Password = v.Password
	case RegisterRequest:
		res.Email = v.Email
		res.Nama = v.Nama
		res.Password = v.Password
	case UpdateRequest:
		res.Nama = v.Nama
		res.Email = v.Email
		res.Password = v.Password
		res.Telepon = v.Telepon
		res.Deskripsi = v.Deskripsi
		res.Pelajaran = v.Pelajaran
		res.Alamat = v.Alamat
		res.Avatar = v.Avatar
		res.Ijazah = v.Ijazah
		res.Latitude = v.Latitude
		res.Longitude = v.Longitude
	default:
		return nil
	}

	return &res
}
