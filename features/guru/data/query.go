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
	gq.db.Raw("SELECT COUNT(*) FROM students WHERE deleted_at IS NULL AND email = ?", newGuru.Email).Scan(&existed)
	if existed >= 1 {
		log.Println("guru account already exist (duplicated)")
		return guru.Core{}, errors.New("guru account already exist (duplicated)")
	}
	newGuru.Role = "siswa"
	cnv := CoreToData(newGuru)
	if err := gq.db.Create(&cnv).Error; err != nil {
		log.Println("register query error", err.Error())
		return guru.Core{}, err
	}
	newGuru.ID = cnv.ID

	return ToCore(cnv), nil
}

// Delete implements guru.GuruData
func (*guruQuery) Delete(id uint) error {
	panic("unimplemented")
}

// Profile implements guru.GuruData
func (*guruQuery) Profile(id uint) (guru.Core, error) {
	panic("unimplemented")
}

// Update implements guru.GuruData
func (*guruQuery) Update(id uint, updateData guru.Core) (guru.Core, error) {
	panic("unimplemented")
}
