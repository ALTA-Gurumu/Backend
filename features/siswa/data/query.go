package data

import (
	"Gurumu/features/siswa"
	"errors"
	"log"

	"gorm.io/gorm"
)

type siswaQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) siswa.SiswaData {
	return &siswaQuery{
		db: db,
	}
}

func (sq *siswaQuery) Register(newStudent siswa.Core) (siswa.Core, error) {
	existed := 0
	sq.db.Raw("SELECT COUNT(*) FROM students WHERE deleted_at IS NULL AND email = ?", newStudent.Email).Scan(&existed)
	if existed >= 1 {
		log.Println("student account already exist (duplicated)")
		return siswa.Core{}, errors.New("student account already exist (duplicated)")
	}
	newStudent.Role = "siswa"
	cnv := CoreToData(newStudent)
	if err := sq.db.Create(&cnv).Error; err != nil {
		log.Println("register query error", err.Error())
		return siswa.Core{}, err
	}
	newStudent.ID = cnv.ID

	return ToCore(cnv), nil
}
func (sq *siswaQuery) Profile(id uint) (siswa.Core, error) {
	return siswa.Core{}, nil
}
func (sq *siswaQuery) Update(id uint, updateData siswa.Core) (siswa.Core, error) {
	return siswa.Core{}, nil
}
func (sq *siswaQuery) Delete(id uint) error {
	return nil
}