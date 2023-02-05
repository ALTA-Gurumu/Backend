package service

import (
	"Gurumu/features/ulasan"
	"Gurumu/helper"
	"errors"
	"strings"
)

type ulasanUseCase struct {
	qry ulasan.UlasanData
}

func New(ud ulasan.UlasanData) ulasan.UlasanService {
	return &ulasanUseCase{
		qry: ud,
	}
}

func (uuc *ulasanUseCase) Add(token interface{}, guruId uint, newUlasan ulasan.Core) error {
	siswaId := helper.ExtractToken(token)
	if siswaId <= 0 {
		return errors.New("id tidak valid")
	}

	err := uuc.qry.Add(uint(siswaId), guruId, newUlasan)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data already exist"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}

	return nil
}

func (uuc *ulasanUseCase) GetAll() ([]ulasan.Core, error) {
	res, err := uuc.qry.GetAll()
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []ulasan.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *ulasanUseCase) GetById(guruId uint) ([]ulasan.Core, error) {
	res, err := uuc.qry.GetById(guruId)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return []ulasan.Core{}, errors.New(msg)
	}

	return res, nil
}
