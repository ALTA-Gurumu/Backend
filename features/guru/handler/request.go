package handler

import "Gurumu/features/guru"

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type UpdateRequest struct {
	Username string `json:"username" form:"username"`
	Fullname string `json:"fullname" form:"fullname"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	City     string `json:"city" form:"city"`
	Address  string `json:"address" form:"address"`
	Phone    string `json:"phone" form:"phone"`
}

func ReqToCore(data interface{}) *guru.Core {
	res := guru.Core{}

	switch data.(type) {
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Username = cnv.Username
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Email = cnv.Email
		res.Username = cnv.Username
		res.Fullname = cnv.Fullname
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.Username = cnv.Username
		res.Fullname = cnv.Fullname
		res.Password = cnv.Password
		res.Email = cnv.Email
		res.City = cnv.City
		res.Address = cnv.Address
		res.Phone = cnv.Phone
	default:
		return nil
	}

	return &res
}
