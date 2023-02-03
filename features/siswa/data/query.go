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
	res := Siswa{}
	if err := sq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get by Id query error", err.Error())
		return siswa.Core{}, err
	}

	return ToCore(res), nil
}
func (sq *siswaQuery) Update(id uint, updateData siswa.Core) error {
	return nil
}
func (sq *siswaQuery) Delete(id uint) error {
	data := Siswa{}

	if err := sq.db.Delete(&data, id).Error; err != nil {
		log.Println("Delete query error", err.Error())
		return err
	}

	return nil
}
