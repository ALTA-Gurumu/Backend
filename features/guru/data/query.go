package data

import (
	"Gurumu/features/guru"
	"Gurumu/features/jadwal/data"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type guruQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) guru.GuruData {
	return &guruQuery{
		db: db,
	}
}

func (gq *guruQuery) Register(newGuru guru.Core) (guru.Core, error) {
	existed := 0
	gq.db.Raw("SELECT COUNT(*) FROM gurus WHERE deleted_at IS NULL AND email = ?", newGuru.Email).Scan(&existed)
	if existed >= 1 {
		log.Println("guru account already exist (duplicated)")
		return guru.Core{}, errors.New("guru account duplicated")
	}
	newGuru.Role = "guru"
	newGuru.Verifikasi = false
	cnv := CoreToData(newGuru)
	if err := gq.db.Create(&cnv).Error; err != nil {
		log.Println("register query error", err.Error())
		return guru.Core{}, errors.New("server error")
	}
	newGuru.ID = cnv.ID

	return ToCore(cnv), nil
}

// Profile implements guru.GuruData
func (gq *guruQuery) GetByID(id uint) (interface{}, error) {
	res := Guru{}
	if err := gq.db.Where("id = ?", id).First(&res).Error; err != nil {
		log.Println("Get by Id query error", err.Error())
		return guru.Core{}, err
	}
	if err := gq.db.Preload("Jadwal").Where("id = ?", id).Find(&res).Error; err != nil {
		log.Println("check data user query error", err.Error())
		return nil, err
	}
	resJadwal := Jadwal{}
	if err := gq.db.Where("id = ?", res.ID).Find(&resJadwal).Error; err != nil {
		log.Println("Get by ID query error", err.Error())
		return nil, err
	}

	result := guru.Core{
		Nama:        res.Nama,
		Email:       res.Email,
		Telepon:     res.Telepon,
		LinkedIn:    res.LinkedIn,
		Gelar:       res.Gelar,
		TentangSaya: res.TentangSaya,
		Pengalaman:  res.Pengalaman,
		LokasiAsal:  res.LokasiAsal,
		MetodeBljr:  res.MetodeBljr,
		Tarif:       res.Tarif,
		Pelajaran:   res.Pelajaran,
		Pendidikan:  res.Pendidikan,
		Avatar:      res.Avatar,
		Ijazah:      res.Ijazah,
		Latitude:    res.Latitude,
		Longitude:   res.Longitude,
	}

	for _, v := range res.Jadwal {
		guru := Guru{}
		if err := gq.db.Where("id = ?", v.ID).Find(&guru).Error; err != nil {
			log.Println("Get by ID query error", err.Error())
			return nil, err
		}

		jadwal := data.GuruJadwal{
			ID:      v.ID,
			Tanggal: v.Tanggal,
			Jam:     v.Jam,
			Status:  v.Status,
		}

		result.Jadwal = append(result.Jadwal, jadwal)
	}
	return result, nil
}

func (gq *guruQuery) GetBeranda(loc string, subj string, limit int, offset int) (int, []guru.Core, error) {

	var guruData []GuruRatingBeranda
	var allGuru int64

	tx := gq.db.Raw("SELECT count(id) FROM gurus WHERE verifikasi = 1 AND deleted_at IS NULL").Find(&allGuru)
	if tx.Error != nil {
		return 0, nil, tx.Error
	}
	query := "SELECT gurus.id, gurus.nama, gurus.lokasi_asal, gurus.tentang_saya, gurus.pelajaran, gurus.avatar, gurus.tarif, COALESCE(AVG(ulasans.penilaian), 0) AS avg_rating FROM gurus LEFT JOIN ulasans ON gurus.id = ulasans.guru_id WHERE gurus.verifikasi = 1"

	var rows *sql.Rows
	var err error
	fmt.Println("allGuru", allGuru)
	if loc != "" && subj == "" {
		query = query + " WHERE gurus.lokasi_asal = ?"
		query = query + " GROUP BY gurus.id ORDER BY gurus.id DESC LIMIT ? OFFSET ?"

		rows, err = gq.db.Raw(query, limit, offset, loc).Rows()
		if err != nil {
			return 0, nil, err
		}

	} else if subj != "" && loc != "" {
		query = query + " WHERE gurus.pelajaran = ? AND gurus.lokasi_asal = ? "
		query = query + " GROUP BY gurus.id ORDER BY gurus.id DESC LIMIT ? OFFSET ?"

		rows, err = gq.db.Raw(query, limit, offset, subj, loc).Rows()
		if err != nil {
			return 0, nil, err
		}
	} else if subj != "" && loc == "" {
		query = query + " WHERE gurus.pelajaran = ?"
		query = query + " GROUP BY gurus.id ORDER BY gurus.id DESC LIMIT ? OFFSET ?"

		rows, err = gq.db.Raw(query, limit, offset, subj).Rows()
		if err != nil {
			return 0, nil, err
		}
	} else if subj == "" && loc == "" {
		query = query + " GROUP BY gurus.id ORDER BY gurus.id DESC LIMIT ? OFFSET ?"

		rows, err = gq.db.Raw(query, limit, offset).Rows()
		if err != nil {
			return 0, nil, err
		}
	}

	defer rows.Close()

	for rows.Next() {
		var guru GuruRatingBeranda
		var avgRating float64
		err = rows.Scan(&guru.ID, &guru.Nama, &guru.LokasiAsal, &guru.TentangSaya, &guru.Pelajaran, &guru.Avatar, &guru.Tarif, &avgRating)
		if err != nil {
			return 0, nil, err
		}
		guru.Penilaian = float32(avgRating)
		guruData = append(guruData, guru)
	}

	return int(allGuru), ListRatingToCore(guruData), nil
}

func (gq *guruQuery) Verifikasi(cekData guru.Core) bool {
	if cekData.LokasiAsal != "" && cekData.Telepon != "" && cekData.Pendidikan != "" && cekData.TentangSaya != "" && cekData.Avatar != "" && cekData.LinkedIn != "" && cekData.Ijazah != "" {
		return true
	}
	return false
}

// Update implements guru.GuruData
func (gq *guruQuery) Update(id uint, updateData guru.Core) error {
	cnv := CoreToData(updateData)
	cnv.Verifikasi = gq.Verifikasi(updateData)

	tx := gq.db.Model(&Guru{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&cnv)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		return errors.New("terjadi kesalahan pada server karena data user atau product tidak ditemukan")
	}

	return nil
}

// Delete implements guru.GuruData
func (gq *guruQuery) Delete(id uint) error {
	qry := gq.db.Delete(&Guru{}, id)
	err := qry.Error

	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("tidak ada data user yang terhapus")
	}

	if err != nil {
		log.Println("delete query error")
		return errors.New("tidak bisa menghapus data")
	}

	return nil
}
