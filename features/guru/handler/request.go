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
	Nama        string `json:"nama" form:"nama"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Telepon     string `json:"telepon" form:"telepon"`
	LinkedIn    string `json:"linkedin" form:"linkedin"`
	Gelar       string `json:"gelar" form:"gelar"`
	TentangSaya string `json:"tentangsaya" form:"tentangsaya"`
	Pengalaman  string `json:"pengalaman" form:"pengalaman"`
	LokasiAsal  string `json:"lokasiasal" form:"lokasiasal"`
	Offline     bool   `json:"offline"`
	Online      bool   `json:"online"`
	Tarif       string `json:"tarif" form:"tarif"`
	Pelajaran   string `json:"pelajaran" form:"pelajaran"`
	Pendidikan  string `json:"pendidikan" form:"pendidikan"`
	Avatar      string `json:"avatar" form:"avatar"`
	Ijazah      string `json:"ijazah" form:"ijazah"`
	Latitude    string `json:"latitude" form:"latitude"`
	Longitude   string `json:"longitude" form:"longitude"`
	Role        string
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
		res.LinkedIn = v.LinkedIn
		res.Gelar = v.Gelar
		res.TentangSaya = v.TentangSaya
		res.Pengalaman = v.Pengalaman
		res.LokasiAsal = v.LokasiAsal
		res.Offline = v.Offline
		res.Online = v.Online
		res.Tarif = v.Tarif
		res.Pelajaran = v.Pelajaran
		res.Pendidikan = v.Pendidikan
		res.Avatar = v.Avatar
		res.Ijazah = v.Ijazah
		res.Latitude = v.Latitude
		res.Longitude = v.Longitude
	default:
		return nil
	}

	return &res
}
