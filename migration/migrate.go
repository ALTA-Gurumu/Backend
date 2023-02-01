package migration

import (
	"Gurumu/features/siswa/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(data.Siswa{})
}
