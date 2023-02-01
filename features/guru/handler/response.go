package handler

import "Gurumu/features/guru"

type GuruReponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func ToResponse(data guru.Core) GuruReponse {
	return GuruReponse{
		Nama:  data.Nama,
		Email: data.Email,
	}
}
