package handler

import (
	"Gurumu/features/guru"
	"Gurumu/features/jadwal/data"
)

type GuruReponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func ToResponse(data guru.Core) GuruReponse {
	return GuruReponse{
		Nama:  data.Nama,
		Email: data.Email,
	}
}

type GuruByIDResp struct {
	Nama        string
	Email       string
	Telepon     string
	LinkedIn    string
	Gelar       string
	TentangSaya string
	Pengalaman  string
	LokasiAsal  string
	MetodeBljr  string
	Tarif       int
	Pelajaran   string
	Pendidikan  string
	Avatar      string
	Ijazah      string
	Latitude    float64
	Longitude   float64
	Jadwal      []data.JadwalNG
}

func GuruByID(data guru.Core) GuruByIDResp {
	return GuruByIDResp{
		Nama:        data.Nama,
		Email:       data.Email,
		Telepon:     data.Telepon,
		LinkedIn:    data.LinkedIn,
		Gelar:       data.Gelar,
		TentangSaya: data.TentangSaya,
		Pengalaman:  data.Pengalaman,
		LokasiAsal:  data.LokasiAsal,
		MetodeBljr:  data.MetodeBljr,
		Tarif:       data.Tarif,
		Pelajaran:   data.Pelajaran,
		Pendidikan:  data.Pendidikan,
		Avatar:      data.Avatar,
		Ijazah:      data.Ijazah,
		Latitude:    data.Latitude,
		Longitude:   data.Longitude,
		Jadwal:      data.Jadwal,
	}
}

type ProfileHomeResp struct {
	ID          uint    `json:"guru_id"`
	Nama        string  `json:"nama"`
	LokasiAsal  string  `json:"alamat"`
	TentangSaya string  `json:"judul"`
	Pelajaran   string  `json:"pelajaran"`
	Avatar      string  `json:"avatar"`
	Tarif       int     `json:"tarif"`
	Penilaian   float32 `json:"penilaian"`
}

func ProfileToResponse(data guru.Core) ProfileHomeResp {
	return ProfileHomeResp{
		ID:          data.ID,
		Nama:        data.Nama,
		LokasiAsal:  data.LokasiAsal,
		TentangSaya: data.TentangSaya,
		Pelajaran:   data.Pelajaran,
		Avatar:      data.Avatar,
		Tarif:       data.Tarif,
		Penilaian:   data.Penilaian,
	}
}

func GetProfileHomeResponse(data []guru.Core) []ProfileHomeResp {
	res := []ProfileHomeResp{}
	for _, v := range data {
		res = append(res, ProfileToResponse(v))
	}
	return res
}
