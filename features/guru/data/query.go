package data

import (
	"Gurumu/features/guru"
	"errors"
	"log"

	"gorm.io/gorm"
)

type guruQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) guru.GuruData {
	return &guruQuery{
		db: db,
	}
}

func (gq *guruQuery) Register(newGuru guru.Core) (guru.Core, error) {
	existed := 0
	gq.db.Raw("SELECT COUNT(*) FROM gurus, siswas WHERE deleted_at IS NULL AND email = ?", newGuru.Email).Scan(&existed)
	if existed >= 1 {
		log.Println("guru account already exist (duplicated)")
		return guru.Core{}, errors.New("guru account already exist (duplicated)")
	}
	newGuru.Role = "guru"
	cnv := CoreToData(newGuru)
	if err := gq.db.Create(&cnv).Error; err != nil {
		log.Println("register query error", err.Error())
		return guru.Core{}, err
	}
	newGuru.ID = cnv.ID

	return ToCore(cnv), nil
}

// Profile implements guru.GuruData
func (gq *guruQuery) GetByID(id uint) (guru.Core, error) {
	res := Guru{}
	query := "SELECT gurus.id, gurus.nama, gurus.email, gurus.alamat, gurus.telepon, gurus.deskripsi, gurus.ijazah, gurus.pelajaran, gurus.avatar FROM gurus WHERE gurus.deleted_at IS NULL AND gurus.id = ?"
	tx := gq.db.Raw(query, id).First(&res)
	if tx.Error != nil {
		return guru.Core{}, tx.Error
	}

	return ToCore(res), nil
}

// Update implements guru.GuruData
func (gq *guruQuery) Update(id uint, updateData guru.Core) error {
	cnv := CoreToData(updateData)
	tx := gq.db.Model(&Guru{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&cnv)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		return errors.New("terjadi kesalahan pada server karena data user atau product tidak ditemukan")
	}

	return nil
}

// Delete implements guru.GuruData
func (gq *guruQuery) Delete(id uint) error {
	qry := gq.db.Delete(&Guru{}, id)
	err := qry.Error

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada data user yang terhapus")
	}

	if err != nil {
		log.Println("delete query error")
		return errors.New("tidak bisa menghapus data")
	}

	return nil
}
