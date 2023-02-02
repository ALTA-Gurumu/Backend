package migration

import (
	_guruData "Gurumu/features/guru/data"
	_jadwaldata "Gurumu/features/jadwal/data"
	"Gurumu/features/siswa/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(data.Siswa{})
	db.AutoMigrate(_guruData.Guru{})
	db.AutoMigrate(_jadwaldata.Jadwal{})
}
