package data

import (
	"Gurumu/features/reservasi"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type reservasiData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservasi.ReservasiData {
	return &reservasiData{
		db: db,
	}
}

func (rd *reservasiData) Add(siswaID uint, newReservasi reservasi.Core) (reservasi.Core, error) {
	data := CoreToData(newReservasi)
	data.SiswaID = siswaID

	detailGuru := Guru{}
	err := rd.db.Where("id = ?", newReservasi.GuruID).First(&detailGuru).Error
	if err != nil {
		log.Println("Get detail guru query error")
		return reservasi.Core{}, err
	}

	err = rd.db.Raw("SELECT id FROM jadwals where tanggal = ? AND jam = ?", newReservasi.Tanggal, newReservasi.Jam).First(&data.JadwalID).Error
	if err != nil {
		log.Println("Get jadwal_id query error")
		return reservasi.Core{}, err
	}
	data.TotalTarif = detailGuru.Tarif
	kodePembayaran := "Gurumu -" + fmt.Sprint(data.SiswaID, data.GuruID, time.Now().Minute())

	midtransResp := helper.CreateReservasiTransaction(kodePembayaran, data.TotalTarif, data.MetodePembayaran)

	if midtransResp.TransactionID != "" {
		data.KodeTransaksi = midtransResp.TransactionID
		data.StatusPembayaran = midtransResp.TransactionStatus
		if data.MetodePembayaran == "transfer_va_permata" {
			data.BankPenerima = "Bank Permata"
			data.NomerVa = midtransResp.PermataVaNumber
		} else {
			data.BankPenerima = midtransResp.VaNumbers[0].Bank
			data.NomerVa = midtransResp.VaNumbers[0].VANumber
		}

	} else {
		return reservasi.Core{}, errors.New("gagal menambahkan pembayaran")
	}

	data.StatusPembayaran = midtransResp.TransactionStatus
	err = rd.db.Create(&data).Error
	if err != nil {
		log.Println("add reservasi query error")
		return reservasi.Core{}, err
	}

	res := ToCore(data)
	res.NamaGuru = detailGuru.Nama
	res.Pelajaran = detailGuru.Pelajaran
	res.TotalTarif = detailGuru.Tarif
	res.AlamatSiswa = newReservasi.AlamatSiswa
	res.TeleponSiswa = newReservasi.TeleponSiswa
	return res, nil

}

func (rd *reservasiData) Mysession(userID uint, role, reservasiStatus string) ([]reservasi.Core, error) {
	var (
		sesiSiswa []SesiSiswa
		sesiGuru  []SesiGuru
		query     string
	)

	switch role {
	case "siswa":
		query = "SELECT r.id, r.guru_id, g.nama AS nama_guru, j.tanggal, j.jam , r.tautan_gmet, r.status FROM reservasis r JOIN gurus g ON r.guru_id = g.id JOIN jadwals j ON r.jadwal_id = j.id WHERE r.siswa_id = ? "

		switch reservasiStatus {
		case "selesai":
			query += "AND r.status = 'selesai'"
			result := rd.db.Raw(query, userID).Find(&sesiSiswa)
			if result.Error != nil {
				return []reservasi.Core{}, result.Error
			}

			if len(sesiSiswa) <= 0 {
				return []reservasi.Core{}, fmt.Errorf("data not found")
			}
			return ToListSesikuSiswa(sesiSiswa), nil
		case "ongoing":
			query += "AND r.status = 'ongoing'"
			result := rd.db.Raw(query, userID).Find(&sesiSiswa)
			if result.Error != nil {
				return []reservasi.Core{}, result.Error
			}

			if len(sesiSiswa) <= 0 {
				return []reservasi.Core{}, fmt.Errorf("not found")
			}

			return ToListSesikuSiswa(sesiSiswa), nil
		default:
			return []reservasi.Core{}, nil
		}
	case "guru":
		query = "SELECT r.id, s.nama AS nama_siswa, j.tanggal, j.jam , r.tautan_gmet, r.status FROM reservasis r JOIN siswas s ON r.siswa_id = s.id JOIN jadwals j ON r.jadwal_id = j.id WHERE r.guru_id = ? "

		switch reservasiStatus {

		case "selesai":
			query += "AND r.status = 'selesai'"
			result := rd.db.Raw(query, userID).Find(&sesiGuru)
			if result.Error != nil {
				return []reservasi.Core{}, result.Error
			}
			fmt.Println("asdlfadslf", sesiGuru)
			if len(sesiGuru) <= 0 {
				return []reservasi.Core{}, fmt.Errorf("not found")
			}
			return ToListSesikuGuru(sesiGuru), nil

		case "ongoing":
			query += "AND r.status = 'ongoing'"
			result := rd.db.Raw(query, userID).Find(&sesiGuru)
			if result.Error != nil {
				return []reservasi.Core{}, result.Error
			}

			if len(sesiGuru) <= 0 {
				return []reservasi.Core{}, fmt.Errorf("not found")
			}
			return ToListSesikuGuru(sesiGuru), nil
		default:
			return []reservasi.Core{}, nil
		}
	default:
		return []reservasi.Core{}, nil
	}
}

func (rd *reservasiData) UpdateDataByTrfID(kode string, updateRes reservasi.Core) error {

	cnv := CoreToData(updateRes)

	tx := rd.db.Model(&Reservasi{}).Where("kode_transaksi = ? ", kode).Updates(&cnv)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected <= 0 {
		return errors.New("terjadi kesalahan pada server karena data user atau product tidak ditemukan")
	}

	return nil
}
func (rd *reservasiData) NotificationTransactionStatus(kodeTransaksi, statusTransaksi string) error {
	reservasiData := Reservasi{}

	err := rd.db.First(&reservasiData, " kode_transaksi = ?", kodeTransaksi).Error
	if err != nil {
		log.Println("transaction not found: ", err.Error())
		return err
	}

	if statusTransaksi == "capture" {
		if statusTransaksi == "challenge" {
			reservasiData.StatusPembayaran = "challenge"
		} else if statusTransaksi == "accept" {
			reservasiData.StatusPembayaran = "success"
		}
	} else if statusTransaksi == "settlement" {
		reservasiData.StatusPembayaran = "success"
	} else if statusTransaksi == "cancel" || statusTransaksi == "expire" {
		reservasiData.StatusPembayaran = "failure"
	} else if statusTransaksi == "pending" {
		reservasiData.StatusPembayaran = "waiting payment"
	} else {
		reservasiData.StatusPembayaran = statusTransaksi
	}

	rowsAffected := rd.db.Save(&reservasiData)
	if rowsAffected.RowsAffected <= 0 {
		log.Println("error update status pembayaran")
		return errors.New("error update status pembayaran")
	}

	if reservasiData.StatusPembayaran == "success" {
		reservasiData.Status = "ongoing"
		aff := rd.db.Save(&reservasiData)
		if aff.RowsAffected <= 0 {
			log.Println("error update status reservasi")
			return errors.New("error update status reservasi")
		}

		detailGuru := Guru{}
		err = rd.db.Where("id = ?", reservasiData.GuruID).First(&detailGuru).Error
		if err != nil {
			log.Println("Get detail guru query error")
			return errors.New("error get detail guru query")
		}

		detailSiswa := Siswa{}
		err = rd.db.Where("id = ?", reservasiData.SiswaID).First(&detailSiswa).Error
		if err != nil {
			log.Println("Get detail siswa query error")
			return errors.New("error get detail siswa query")
		}

		detailJadwal := Jadwal{}
		err = rd.db.Where("id = ?", reservasiData.JadwalID).First(&detailJadwal).Error
		if err != nil {
			log.Println("Get detail jadwal query error")
			return errors.New("error get jadwal guru query")
		}

		res, err := helper.CalendarEvent(detailGuru.Email, detailSiswa.Email, detailGuru.Pelajaran, detailJadwal.Tanggal, detailJadwal.Jam)
		if err != nil {
			log.Println("Create event query error")
			return err
		}
		fmt.Println(res)
	}
	return nil
}
