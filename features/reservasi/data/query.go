package data

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type reservasiData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservasi.ReservasiData {
	return &reservasiData{
		db: db,
	}
}

func (rd *reservasiData) Add(siswaID uint, newReservasi reservasi.Core) (reservasi.Core, error) {
	data := CoreToData(newReservasi)
	data.SiswaID = siswaID

	detailGuru := Guru{}
	err := rd.db.Where("id = ?", newReservasi.GuruID).First(&detailGuru).Error
	if err != nil {
		log.Println("Get detail guru query error")
		return reservasi.Core{}, err
	}

	err = rd.db.Raw("SELECT id FROM jadwals where tanggal = ? AND jam = ?", newReservasi.Tanggal, newReservasi.Jam).First(&data.JadwalID).Error
	if err != nil {
		log.Println("Get jadwal_id query error")
		return reservasi.Core{}, err
	}
	data.TotalTarif = detailGuru.Tarif
	kodePembayaran := "Gurumu -" + fmt.Sprint(data.SiswaID) + fmt.Sprint(data.GuruID)

	midtransResp := helper.CreateReservasiTransaction(kodePembayaran, data.TotalTarif, data.MetodePembayaran)

	if midtransResp.TransactionID != "" {
		data.KodeTransaksi = midtransResp.TransactionID
		data.StatusPembayaran = midtransResp.TransactionStatus
		if data.MetodePembayaran == "transfer_va_permata" {
			data.BankPenerima = "Bank Permata"
			data.NomerVa = midtransResp.PermataVaNumber
		} else {
			data.BankPenerima = midtransResp.VaNumbers[0].Bank
			data.NomerVa = midtransResp.VaNumbers[0].VANumber
		}

	} else {
		return reservasi.Core{}, errors.New("gagal menambahkan pembayaran")
	}

	data.Status = "Belum melakukan Pembayaran"

	tautanGmet, err := helper.Calendar(detailGuru.Email, newReservasi.Tanggal, newReservasi.AlamatSiswa)
	if err != nil {
		fmt.Println("gagal menambahkan ke kalender")
		return reservasi.Core{}, errors.New("gagal menambahkan ke kalender")
	}

	fmt.Println(tautanGmet)
	data.TautanGmet = tautanGmet
	err = rd.db.Create(&data).Error
	if err != nil {
		log.Println("add reservasi query error")
		return reservasi.Core{}, err
	}

	res := ToCore(data)
	res.NamaGuru = detailGuru.Nama
	res.Pelajaran = detailGuru.Pelajaran
	res.TotalTarif = detailGuru.Tarif
	res.AlamatSiswa = newReservasi.AlamatSiswa
	res.TeleponSiswa = newReservasi.TeleponSiswa

	return res, nil

}

func (rd *reservasiData) Mysession(userID uint, role, reservasiStatus string) ([]reservasi.Core, error) {
	var sesiSiswa = []SesiSiswa{}
	var sesiGuru = []SesiGuru{}

	if role == "siswa" {
		query := "SELECT r.id, g.nama, j.tanggal, j.jam , r.tautan_gmet, r.status FROM reservasis r JOIN gurus g ON r.guru_id = g.id JOIN jadwals j ON r.jadwal_id = j.id WHERE r.siswa_id = ? "
		if reservasiStatus == "selesai" {
			query += "AND r.status = ?"

			err := rd.db.Raw(query, userID, reservasiStatus).Find(&sesiSiswa).Error
			if err != nil {
				log.Println("get sesi siswa query error")
				return []reservasi.Core{}, err
			}

			return ToListSesikuSiswa(sesiSiswa), nil

		} else if reservasiStatus == "ongoing" {
			query += "AND r.status = ?"

			err := rd.db.Raw(query, userID, reservasiStatus).Find(&sesiSiswa).Error
			if err != nil {
				return []reservasi.Core{}, err
			}
			return ToListSesikuSiswa(sesiSiswa), nil
		}

	} else if role == "guru" {
		query := "SELECT r.id, s.nama, j.tanggal, j.jam , r.tautan_gmet, r.status FROM reservasis r JOIN siswas s ON r.guru_id = s.id JOIN jadwals j ON r.jadwal_id = j.id WHERE r.guru_id = ? "
		if reservasiStatus == "selesai" {
			query += "AND r.status = ?"

			err := rd.db.Raw(query, userID, reservasiStatus).Find(&sesiGuru).Error
			if err != nil {
				log.Println("get sesi guru query error")
				return []reservasi.Core{}, err
			}

			return ToListSesikuGuru(sesiGuru), nil

		} else if reservasiStatus == "ongoing" {
			query += "AND r.status = ?"

			err := rd.db.Raw(query, userID, reservasiStatus).Find(&sesiGuru).Error
			if err != nil {
				return []reservasi.Core{}, err
			}
			return ToListSesikuGuru(sesiGuru), nil

		}

	}
	return []reservasi.Core{}, nil
}
