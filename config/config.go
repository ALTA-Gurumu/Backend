package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var (
	JWT_KEY           string = ""
	KEYID             string = ""
	ACCESSKEY         string = ""
	MIDTRANSSERVERKEY string = ""
)

type AppConfig struct {
	DBUser            string
	DBPass            string
	DBHost            string
	DBPort            int
	DBName            string
	jwtKey            string
	keyid             string
	accesskey         string
	midtransserverkey string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("MIDTRANSSERVERKEY"); found {
		app.midtransserverkey = val
		isRead = false
		MIDTRANSSERVERKEY = val
	}
	if val, found := os.LookupEnv("KEYID"); found {
		app.keyid = val
		isRead = false
		KEYID = val
	}
	if val, found := os.LookupEnv("ACCESSKEY"); found {
		app.accesskey = val
		isRead = false
		ACCESSKEY = val

	}
	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
		JWT_KEY = val
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		err := godotenv.Load("local.env")
		if err != nil {
			fmt.Println("Error saat baca env", err.Error())
			return nil
		}

		app.DBUser = os.Getenv("DBUSER")
		app.DBPass = os.Getenv("DBPASS")
		app.DBHost = os.Getenv("DBHOST")
		readData := os.Getenv("DBPORT")
		app.DBPort, err = strconv.Atoi(readData)
		if err != nil {
			fmt.Println("Error saat convert", err.Error())
			return nil
		}
		app.DBName = os.Getenv("DBNAME")
		app.jwtKey = os.Getenv("JWTKEY")
		app.keyid = os.Getenv("KEYID")
		app.accesskey = os.Getenv("ACCESSKEY")
		app.midtransserverkey = os.Getenv("MIDTRANSSERVERKEY")

		JWT_KEY = app.jwtKey
		KEYID = app.keyid
		ACCESSKEY = app.accesskey
		MIDTRANSSERVERKEY = app.midtransserverkey
	}

	return &app
}

func S3Config() *session.Session {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(KEYID, ACCESSKEY, ""),
	}
	s3Session, _ := session.NewSession(s3Config)
	return s3Session
}

func MidtransSnapClient() snap.Client {
	s := snap.Client{}
	s.New(MIDTRANSSERVERKEY, midtrans.Sandbox)
	return s
}
