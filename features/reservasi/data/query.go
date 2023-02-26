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

	existingReservation := Reservasi{}
	err = rd.db.Where("siswa_id = ? AND jadwal_id = ?", data.SiswaID, data.JadwalID).First(&existingReservation).Error
	if err == nil {

		return reservasi.Core{}, errors.New("reservasi sudah ditambahkan sebelumnya")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {

		log.Println("Get existing reservations query error")
		return reservasi.Core{}, err
	}

	data.TotalTarif = detailGuru.Tarif
	kodePembayaran := "GRM/" + fmt.Sprint(data.SiswaID) + fmt.Sprint(data.GuruID) + fmt.Sprint(time.Now().Hour()) + fmt.Sprint(time.Now().Minute())

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
		return reservasi.Core{}, errors.New("gagal menambahkan pembayaran, cobalah beberapa saat lagi")
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
				return []reservasi.Core{}, errors.New("data not found")
			}
			return ToListSesikuSiswa(sesiSiswa), nil
		case "ongoing":
			query += "AND r.status = 'ongoing'"
			result := rd.db.Raw(query, userID).Find(&sesiSiswa)
			if result.Error != nil {
				return []reservasi.Core{}, result.Error
			}

			if len(sesiSiswa) <= 0 {
				return []reservasi.Core{}, errors.New("not found")
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

			if len(sesiGuru) <= 0 {
				return []reservasi.Core{}, errors.New("tidak ditemukan")
			}
			return ToListSesikuGuru(sesiGuru), nil

		case "ongoing":
			query += "AND r.status = 'ongoing'"
			result := rd.db.Raw(query, userID).Find(&sesiGuru)
			if result.Error != nil {
				return []reservasi.Core{}, result.Error
			}

			if len(sesiGuru) <= 0 {
				return []reservasi.Core{}, errors.New("tidak ditemukan")
			}
			return ToListSesikuGuru(sesiGuru), nil
		default:
			return []reservasi.Core{}, nil
		}
	default:
		return []reservasi.Core{}, nil
	}
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
			return errors.New("gagal update status reservasi")
		}

		detailGuru := Guru{}
		err = rd.db.Where("id = ?", reservasiData.GuruID).First(&detailGuru).Error
		if err != nil {
			log.Println("Get detail guru query error")
			return errors.New("data guru tidak ditemukan")
		}

		detailSiswa := Siswa{}
		err = rd.db.Where("id = ?", reservasiData.SiswaID).First(&detailSiswa).Error
		if err != nil {
			log.Println("Get detail siswa query error")
			return errors.New("data siswa tidak ditemukan")
		}

		detailJadwal := Jadwal{}
		err = rd.db.Where("id = ?", reservasiData.JadwalID).First(&detailJadwal).Error
		if err != nil {
			log.Println("Get detail jadwal query error")
			return errors.New("data jadwal tidak ditemukan")
		}

		// 	layout := "2006-01-02 15:04:05"
		// 	value := detailJadwal.Tanggal + " " + detailJadwal.Jam + ":00"
		// 	dateTime, err := time.Parse(layout, value)
		// 	if err != nil {
		// 		return errors.New("failed convert datetime")
		// 	}
		// 	fmt.Println(detailGuru.Email, detailSiswa.Email)
		// 	tautanGmeet := helper.CreateEvent(
		// 		&calendar.Event{
		// 			Summary:     "Gurumu - Kelas " + detailGuru.Pelajaran + " anda",
		// 			Location:    "",
		// 			Description: "Kelas akan berlangsung pada " + detailJadwal.Tanggal + " pada " + detailJadwal.Jam + ". Harap datang tepat waktu dan pastikan untuk bergabung dengan panggilan video tepat waktu.",
		// 			ConferenceData: &calendar.ConferenceData{
		// 				CreateRequest: &calendar.CreateConferenceRequest{
		// 					RequestId: "sfsfs",
		// 					ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
		// 						Type: "hangoutsMeet"},
		// 					Status: &calendar.ConferenceRequestStatus{
		// 						StatusCode: "success"},
		// 				}},

		// 			Start: &calendar.EventDateTime{
		// 				DateTime: dateTime.Format(time.RFC3339),
		// 				TimeZone: "Asia/Jakarta",
		// 			},
		// 			End: &calendar.EventDateTime{
		// 				DateTime: dateTime.Add(time.Hour * 1).Format(time.RFC3339),
		// 				TimeZone: "Asia/Jakarta",
		// 			},

		// 			Attendees: []*calendar.EventAttendee{
		// 				{Email: detailGuru.Email},
		// 				{Email: detailSiswa.Email},
		// 			},
		// 			Reminders: &calendar.EventReminders{
		// 				UseDefault: true,
		// 				// Overrides: []*calendar.EventReminder{
		// 				// 	{Method: "email", Minutes: 10},
		// 				// },
		// 			},
		// 		})

		reservasiData.TautanGmet = "-"
		aff = rd.db.Save(&reservasiData)
		if aff.RowsAffected <= 0 {
			log.Println("error update tautan gmeet reservasi")
			return errors.New("gagal update tautan gmeet")
		}

		jadwalData := Jadwal{}

		err = rd.db.First(&jadwalData, "id = ?", reservasiData.JadwalID).Error
		if err != nil {
			log.Println("jadwal not found")
			return err
		}
		fmt.Println(jadwalData)

		jadwalData.Status = "Telah direservasi"
		aff = rd.db.Save(&jadwalData)
		if aff.RowsAffected <= 0 {
			log.Println("error update status jadwal")
			return errors.New("gagal update status jadwal")
		}

	}
	return nil
}
func (rd *reservasiData) UpdateStatus(userID uint, reservasiID uint) error {
	reservasiData := Reservasi{}

	err := rd.db.First(&reservasiData, " id = ?", reservasiID).Error
	if err != nil {
		log.Println("resevasi not found: ", err.Error())
		return errors.New("data reservasi tidak ditemukan")
	}

	reservasiData.Status = "selesai"
	aff := rd.db.Save(&reservasiData)
	if aff.RowsAffected <= 0 {
		log.Println("error update status reservasi")
		return errors.New("gagal update status reservasi")
	}
	return nil

}
