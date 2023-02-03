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
	gq.db.Raw("SELECT COUNT(*) FROM gurus WHERE deleted_at IS NULL AND email = ?", newGuru.Email).Scan(&existed)
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
func (gq *guruQuery) Profile(id uint) (guru.Core, error) {
	
}

// Update implements guru.GuruData
func (gq *guruQuery) Update(id uint, updateData guru.Core) (guru.Core, error) {
	cnv := CoreToData(updateData)
	tx := gq.db.Model(&Guru{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&cnv)
	if tx.Error != nil {
		return tx.Error
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
