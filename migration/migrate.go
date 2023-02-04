package migration

import (
	_guruData "Gurumu/features/guru/data"
	_jadwalData "Gurumu/features/jadwal/data"
	_reservasiData "Gurumu/features/reservasi/data"
	"Gurumu/features/siswa/data"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(data.Siswa{})
	db.AutoMigrate(_guruData.Guru{})
	db.AutoMigrate(_jadwalData.Jadwal{})
	db.AutoMigrate(_reservasiData.Reservasi{})

}
