package data

import (
	"Gurumu/features/ulasan"

	"gorm.io/gorm"
)

// type Guru struct {
// 	gorm.Model
// 	Nama        string
// 	Email       string
// 	Password    string
// 	Telepon     string
// 	LinkedIn    string
// 	Gelar       string
// 	TentangSaya string
// 	Pengalaman  string
// 	LokasiAsal  string
// 	Offline     bool
// 	Online      bool
// 	Tarif       string
// 	Pelajaran   string
// 	Pendidikan  string
// 	Avatar      string
// 	Ijazah      string
// 	Role        string
// 	Latitude    string
// 	Longitude   string
// 	Ulasan      []Ulasan `gorm:"foreignKey:GuruId;references:ID"`
// }

// type Siswa struct {
// 	gorm.Model
// 	Email    string
// 	Password string
// 	Nama     string
// 	Telepon  string
// 	Alamat   string
// 	Avatar   string
// 	Role     string
// 	Ulasan   []Ulasan `gorm:"foreignKey:SiswaId;references:ID"`
// }

type Ulasan struct {
	gorm.Model
	GuruId    uint
	SiswaId   uint
	Ulasan    string
	Penilaian float32
}

func CoreToData(core ulasan.Core) Ulasan {
	return Ulasan{
		Model:     gorm.Model{ID: core.ID},
		GuruId:    core.GuruId,
		SiswaId:   core.SiswaId,
		Ulasan:    core.Ulasan,
		Penilaian: core.Penilaian,
	}
}

func DataToCore(data Ulasan) ulasan.Core {
	return ulasan.Core{
		ID:        data.ID,
		GuruId:    data.GuruId,
		SiswaId:   data.SiswaId,
		Ulasan:    data.Ulasan,
		Penilaian: data.Penilaian,
	}
}

// Untuk Ulasan tiap Guru
type UlasanGuru struct {
	ID        uint
	NamaSiswa string
	Penilaian float32
	Ulasan    string
}

func (dataModel *UlasanGuru) ModelsToCore() ulasan.Core {
	return ulasan.Core{
		ID:        dataModel.ID,
		NamaSiswa: dataModel.NamaSiswa,
		Penilaian: dataModel.Penilaian,
		Ulasan:    dataModel.Ulasan,
	}
}

func ListModelsToCore(dataModels []UlasanGuru) []ulasan.Core {
	var dataCore []ulasan.Core
	for _, val := range dataModels {
		dataCore = append(dataCore, val.ModelsToCore())
	}
	return dataCore
}

// Untuk Semua ulasan
type AllUlasan struct {
	ID        uint
	GuruId    uint
	NamaGuru  string
	Penilaian float32
	Ulasan    string
}

func (dataModel *AllUlasan) AllModelsToCore() ulasan.Core {
	return ulasan.Core{
		ID:        dataModel.ID,
		GuruId:    dataModel.GuruId,
		NamaGuru:  dataModel.NamaGuru,
		Penilaian: dataModel.Penilaian,
		Ulasan:    dataModel.Ulasan,
	}
}

func ListAllModelsToCore(dataModels []AllUlasan) []ulasan.Core {
	var dataCore []ulasan.Core
	for _, val := range dataModels {
		dataCore = append(dataCore, val.AllModelsToCore())
	}
	return dataCore
}
