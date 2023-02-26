package data

import (
	"Gurumu/features/jadwal"
	"errors"
	"log"

	"gorm.io/gorm"
)

type jadwalQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) jadwal.JadwalData {
	return &jadwalQuery{
		db: db,
	}
}

func (jq *jadwalQuery) Add(guruID uint, newJadwal jadwal.Core) (jadwal.Core, error) {
	data := CoreToData(newJadwal)

	existedJadwal := Jadwal{}
	err := jq.db.Where("id = ? AND tanggal = ? AND jam = ?", guruID, data.Tanggal, data.Jam).Find(&existedJadwal).Error
	if err != nil {
		return jadwal.Core{}, err
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return jadwal.Core{}, err
	}

	data.GuruID = guruID
	data.Status = "Tersedia"
	err = jq.db.Create(&data).Error
	if err != nil {
		log.Println("error saat query tambah jadwal")
		return jadwal.Core{}, err
	}

	newJadwal.ID = data.ID
	newJadwal.GuruID = data.GuruID

	return newJadwal, nil
}

func (jq *jadwalQuery) GetJadwal(guruID uint) ([]jadwal.Core, error) {
	listJadwal := []Jadwal{}

	err := jq.db.Where("guru_id = ?", guruID).Find(&listJadwal).Error
	if err != nil {
		log.Println("list jadwal query error", err.Error())
		return []jadwal.Core{}, err
	}

	return ListToCore(listJadwal), nil
}
