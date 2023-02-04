package data

import (
	"Gurumu/features/reservasi"
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

	err = rd.db.Raw("SELECT id FROM jadwals where tanggal = ? AND jam = ?", newReservasi.Tanggal, newReservasi.Jam).First(&data.JadwalID).Error
	if err != nil {
		log.Println("Get jadwal_id query error")
		return reservasi.Core{}, err
	}

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
