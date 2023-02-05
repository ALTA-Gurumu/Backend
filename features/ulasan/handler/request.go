package handler

import "Gurumu/features/ulasan"

type UlasanRegisterRequest struct {
	Ulasan    string  `json:"ulasan" form:"ulasan"`
	Penilaian float32 `json:"penilaian" form:"penilaian"`
}

func ToCore(data interface{}) *ulasan.Core {
	res := ulasan.Core{}

	switch data.(type) {
	case UlasanRegisterRequest:
		cnv := data.(UlasanRegisterRequest)
		res.Ulasan = cnv.Ulasan
		res.Penilaian = cnv.Penilaian
	default:
		return nil
	}
	return &res
}
