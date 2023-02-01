package handler

import "Gurumu/features/guru"

type UserReponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
}

func ToResponse(data guru.Core) UserReponse {
	return UserReponse{
		Nama:  data.Nama,
		Email: data.Email,
	}
}
