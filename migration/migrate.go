package migration

import (
	_guruData "Gurumu/features/guru/data"
	"Gurumu/features/siswa/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(data.Siswa{})
	db.AutoMigrate(_guruData.Guru{})
}
