package handler

import "Gurumu/features/reservasi"

type AddReservasiResponse struct {
	ID               uint   `json:"id" form:"id"`
	NamaGuru         string `json:"nama_guru" form:"nama_guru"`
	MetodeBelajar    string `json:"metode_belajar" form:"metode_belajar"`
	Pelajaran        string `json:"pelajaran" form:"pelajaran"`
	TotalTarif       int    `json:"total_tarif" form:"total_tarif"`
	AlamatSiswa      string `json:"alamat_siswa" form:"alamat_siswa"`
	TeleponSiswa     string `json:"telepon_siswa" form:"telepon_siswa"`
	KodeTransaksi    string `json:"kode_transaksi" form:"kode_transaksi"`
	MetodePembayaran string `json:"metode_pembayaran" form:"metode_pembayaran"`
	NomerVa          string `json:"nomer_va" form:"nomer_va"`
	KodeQr           string `json:"kode_qr" form:"kode_qr"`
	BankPenerima     string `json:"bank_penerima" form:"bank_penerima"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
	TautanGmet       string `json:"tautan_gmet" form:"tautan_gmet"`
	Status           string `json:"status" form:"status"`
}

type GetSesikuGuruResponse struct {
	ID         uint   `json:"reservasi_id" form:"reservasi_id"`
	NamaSiswa  string `json:"nama_siswa" form:"nama_siswa"`
	Tanggal    string `json:"tanggal" form:"tanggal"`
	Jam        string `json:"jam" form:"jam"`
	TautanGmet string `json:"tautan_gmet" form:"tautan_gmet"`
	Status     string `json:"status" form:"status"`
}

type GetSesikuSiswaResponse struct {
	ID         uint   `json:"reservasi_id" form:"reservasi_id"`
	NamaGuru   string `json:"nama_guru" form:"nama_guru"`
	Tanggal    string `json:"tanggal" form:"tanggal"`
	Jam        string `json:"jam" form:"jam"`
	TautanGmet string `json:"tautan_gmet" form:"tautan_gmet"`
	Status     string `json:"status" form:"status"`
}

func ToAddReservasiResponse(data reservasi.Core) AddReservasiResponse {
	return AddReservasiResponse{
		ID:               data.ID,
		NamaGuru:         data.NamaGuru,
		MetodeBelajar:    data.MetodeBelajar,
		Pelajaran:        data.Pelajaran,
		TotalTarif:       data.TotalTarif,
		AlamatSiswa:      data.AlamatSiswa,
		TeleponSiswa:     data.TeleponSiswa,
		KodeTransaksi:    data.KodeTransaksi,
		MetodePembayaran: data.MetodePembayaran,
		NomerVa:          data.NomerVa,
		KodeQr:           data.KodeQr,
		BankPenerima:     data.BankPenerima,
		StatusPembayaran: data.StatusPembayaran,
		TautanGmet:       data.TautanGmet,
		Status:           data.Status,
	}
}

func ToSesikuGuruResponse(data reservasi.Core) GetSesikuGuruResponse {
	return GetSesikuGuruResponse{
		ID:         data.ID,
		NamaSiswa:  data.NamaSiswa,
		Tanggal:    data.Tanggal,
		Jam:        data.Jam,
		TautanGmet: data.TautanGmet,
		Status:     data.Status,
	}
}
func ToSesikuSiswaResponse(data reservasi.Core) GetSesikuSiswaResponse {
	return GetSesikuSiswaResponse{
		ID:         data.ID,
		NamaGuru:   data.NamaGuru,
		Tanggal:    data.Tanggal,
		Jam:        data.Jam,
		TautanGmet: data.TautanGmet,
		Status:     data.Status,
	}
}

func ToListSesikuGuruResponse(data []reservasi.Core) []GetSesikuGuruResponse {
	var listSesiGuru = []GetSesikuGuruResponse{}
	for _, sesiGuru := range data {
		listSesiGuru = append(listSesiGuru, ToSesikuGuruResponse(sesiGuru))
	}

	return listSesiGuru
}

func ToListSesikuSiswaResponse(data []reservasi.Core) []GetSesikuSiswaResponse {
	var listSesiSiswa = []GetSesikuSiswaResponse{}
	for _, sesiSiswa := range data {
		listSesiSiswa = append(listSesiSiswa, ToSesikuSiswaResponse(sesiSiswa))
	}

	return listSesiSiswa
}
