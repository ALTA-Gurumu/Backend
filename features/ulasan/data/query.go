package data

import (
	"Gurumu/features/ulasan"
	"log"

	"gorm.io/gorm"
)

type ulasanQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) ulasan.UlasanData {
	return &ulasanQuery{
		db: db,
	}
}

func (uq *ulasanQuery) Add(siswaId, guruId uint, newUlasan ulasan.Core) error {
	cnv := CoreToData(newUlasan)
	cnv.GuruId = guruId
	cnv.SiswaId = siswaId
	if err := uq.db.Create(&cnv).Error; err != nil {
		log.Println("register ulasan query error", err.Error())
		return err
	}

	return nil
}
func (uq *ulasanQuery) GetAll() ([]ulasan.Core, error) {
	res := []AllUlasan{}
	err := uq.db.Raw("SELECT ulasans.id, gurus.nama as nama_guru, ulasans.penilaian, ulasans.ulasan FROM ulasans JOIN gurus ON ulasans.guru_id = gurus.id WHERE ulasans.deleted_at is NULL").Find(&res).Error
	if err != nil {
		log.Println("all ulasan query error")
		return []ulasan.Core{}, err
	}

	return ListAllModelsToCore(res), nil
}

func (uq *ulasanQuery) GetById(guruId uint) ([]ulasan.Core, error) {
	res := []UlasanGuru{}
	err := uq.db.Raw("SELECT ulasans.id, siswas.nama as nama_siswa, ulasans.penilaian, ulasans.ulasan FROM ulasans JOIN siswas ON ulasans.siswa_id = siswas.id WHERE ulasans.deleted_at is NULL AND ulasans.guru_id = ?", guruId).Find(&res).Error
	if err != nil {
		log.Println("get ulasan by id guru query error")
		return []ulasan.Core{}, err
	}

	return ListModelsToCore(res), nil
}