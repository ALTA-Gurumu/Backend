package data

import (
	"Gurumu/features/autentikasi"

	"gorm.io/gorm"
)

type Siswa struct {
	gorm.Model
	Email    string
	Password string
	Nama     string
	Telepon  string
	Alamat   string
	Avatar   string
	Role     string
}
type Guru struct {
	gorm.Model
	Email    string
	Password string
	Nama     string
	Telepon  string
	Alamat   string
	Avatar   string
	Role     string
}

func SiswaToCore(data Siswa) autentikasi.Core {
	return autentikasi.Core{
		ID:       data.ID,
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}

func GuruToCore(data Guru) autentikasi.Core {
	return autentikasi.Core{
		ID:       data.ID,
		Nama:     data.Nama,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	}
}
