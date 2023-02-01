package handler

import "Gurumu/features/autentikasi"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(data LoginRequest) *autentikasi.Core {
	return &autentikasi.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}
