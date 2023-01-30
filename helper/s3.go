package helper

import (
	"Gurumu/config"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var ObjectURL string = "https://ecommercegroup7.s3.ap-southeast-1.amazonaws.com/"

func UploadProfilePhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ecommercegroup7"),
		Key:    aws.String("files/user/" + email + "/profile-photo" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload profile photo")
	}
	path := ObjectURL + "files/user/" + email + "/profile-photo" + ext
	return path, nil
}

func UploadProductImageS3(file multipart.FileHeader, userID int) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	cnv := strconv.Itoa(userID)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ecommercegroup7"),
		Key:    aws.String("files/product/" + cnv + "/" + file.Filename),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload product image")
	}
	path := ObjectURL + "files/product/" + cnv + "/" + file.Filename
	return path, nil
}
