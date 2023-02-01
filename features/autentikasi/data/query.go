package data

import (
	"Gurumu/features/autentikasi"
	"errors"
	"log"

	"gorm.io/gorm"
)

type autentikasiQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) autentikasi.AutentikasiData {
	return &autentikasiQuery{
		db: db,
	}
}

func (aq *autentikasiQuery) Login(email string) (autentikasi.Core, error) {
	//Cek Akun Siswa
	var cekData = Siswa{}
	if err := aq.db.Where("email = ?", email).First(&cekData).Error; err != nil {
		//Cek Akun Guru
		var cekData = Guru{}
		if err := aq.db.Where("email = ?", email).First(&cekData).Error; err != nil {
			log.Println("login query error", err.Error())
			return autentikasi.Core{}, errors.New("data tidak ditemukan")
		}
		return GuruToCore(cekData), nil

	}

	return SiswaToCore(cekData), nil

}
