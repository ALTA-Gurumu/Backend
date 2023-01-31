package helper

import (
	"Gurumu/config"
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ObjectURL string = "https://capstonegurumu.s3.ap-southeast-1.amazonaws.com/"

func UploadStudentProfilePhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("capstonegurumu"),
		Key:    aws.String("files/students/" + email + "/profile-photo" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload student's profile photo")
	}
	path := ObjectURL + "files/students/" + email + "/profile-photo" + ext
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
		Bucket: aws.String("capstonegurumu"),
		Key:    aws.String("files/teachers/" + email + "/profile-photo" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload teacher's profile photo")
	}
	path := ObjectURL + "files/teachers/" + email + "/profile-photo" + ext
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
		Bucket: aws.String("capstonegurumu"),
		Key:    aws.String("files/teachers/" + email + "/certificate" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload teacher's certificate")
	}
	path := ObjectURL + "files/teachers/" + email + "/certificate" + ext
	return path, nil
}
