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
	Offline     string
	Online      string
	Tarif       string
	Pelajaran   string
	Pendidikan  string
	Avatar      string
	Ijazah      string
	Latitude    string
	Longitude   string
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
	ID          uint
	Nama        string
	LokasiAsal  string
	TentangSaya string
	Pelajaran   string
	Avatar      string
	Penilaian   float32
}

func ProfileToResponse(data guru.Core) ProfileHomeResp {
	return ProfileHomeResp{
		ID:          data.ID,
		Nama:        data.Nama,
		LokasiAsal:  data.LokasiAsal,
		TentangSaya: data.TentangSaya,
		Pelajaran:   data.Pelajaran,
		Avatar:      data.Avatar,
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
