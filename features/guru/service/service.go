package service

import (
	"Gurumu/features/guru"
	"mime/multipart"

	"github.com/go-playground/validator"
)

type guruUseCase struct {
	qry guru.GuruData
	vld *validator.Validate
}

func New(gd guru.GuruData) guru.GuruService {
	return &guruUseCase{
		qry: gd,
		vld: validator.New(),
	}
}

// Register implements guru.GuruService
func (*guruUseCase) Register(newGuru guru.Core) (guru.Core, error) {
	panic("unimplemented")
}

// Delete implements guru.GuruService
func (*guruUseCase) Delete(token interface{}) error {
	panic("unimplemented")
}

// Profile implements guru.GuruService
func (*guruUseCase) Profile(token interface{}) (guru.Core, error) {
	panic("unimplemented")
}

// Update implements guru.GuruService
func (*guruUseCase) Update(token interface{}, updateData guru.Core, avatar *multipart.FileHeader) (guru.Core, error) {
	panic("unimplemented")
}
