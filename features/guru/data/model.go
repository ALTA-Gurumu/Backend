package data

import (
	"Gurumu/features/guru"

	"gorm.io/gorm"
)

type Guru struct {
	gorm.Model
	Email     string
	Password  string
	Nama      string
	Telepon   string
	Deskripsi string
	Ijazah    string
	Pelajaran string
	Alamat    string
	Avatar    string
	Role      string
}

func ToCore(data Guru) guru.Core {
	return guru.Core{
		ID:        data.ID,
		Email:     data.Email,
		Password:  data.Password,
		Nama:      data.Nama,
		Telepon:   data.Telepon,
		Deskripsi: data.Deskripsi,
		Ijazah:    data.Ijazah,
		Pelajaran: data.Pelajaran,
		Alamat:    data.Alamat,
		Avatar:    data.Avatar,
		Role:      data.Role,
	}
}

func CoreToData(core guru.Core) Guru {
	return Guru{
		Model:     gorm.Model{ID: core.ID},
		Email:     core.Email,
		Password:  core.Password,
		Nama:      core.Nama,
		Telepon:   core.Telepon,
		Deskripsi: core.Deskripsi,
		Ijazah:    core.Ijazah,
		Pelajaran: core.Pelajaran,
		Alamat:    core.Alamat,
		Avatar:    core.Avatar,
		Role:      core.Role,
	}
}
