package data

import (
	"Gurumu/features/siswa"

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

func ToCore(data Siswa) siswa.Core {
	return siswa.Core{
		ID:       data.ID,
		Email:    data.Email,
		Password: data.Password,
		Nama:     data.Nama,
		Telepon:  data.Telepon,
		Alamat:   data.Alamat,
		Avatar:   data.Avatar,
		Role:     data.Role,
	}
}

func CoreToData(core siswa.Core) Siswa {
	return Siswa{
		Model:    gorm.Model{ID: core.ID},
		Email:    core.Email,
		Password: core.Password,
		Nama:     core.Nama,
		Telepon:  core.Telepon,
		Alamat:   core.Alamat,
		Avatar:   core.Avatar,
		Role:     core.Role,
	}
}
