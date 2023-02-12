package helper

import (
	"Gurumu/config"
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader interface {
	UploadStudentProfilePhotoS3(file multipart.FileHeader, email string) (string, error)
	UploadTeacherProfilePhotoS3(file *multipart.FileHeader, email string) (string, error)
	UploadTeacherCertificateS3(file *multipart.FileHeader, email string) (string, error)
}

var ObjectURL string = "https://try123ok.s3.ap-southeast-1.amazonaws.com/"

func UploadStudentProfilePhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("try123ok"),
		Key:    aws.String("files/siswa/" + email + "/" + file.Filename),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload avatar siswa")
	}
	path := ObjectURL + "files/siswa/" + email + "/" + file.Filename
	return path, nil
}

func UploadTeacherProfilePhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("try123ok"),
		Key:    aws.String("files/guru/" + email + "/avatar" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload avatar guru")
	}
	path := ObjectURL + "files/guru/" + email + "/avatar" + ext
	return path, nil
}

func UploadTeacherCertificateS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("try123ok"),
		Key:    aws.String("files/guru/" + email + "/certificate" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload certificate guru")
	}
	path := ObjectURL + "files/guru/" + email + "/certificate" + ext
	return path, nil
}
