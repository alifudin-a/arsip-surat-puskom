package helper

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/jlaffaye/ftp"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

// Upload: file to server
func Upload(file []byte, filename string) (err error) {
	logrus.Info("=================> Processing upload file to server")
	conn, err := connectFTP()

	if err != nil {
		logrus.Error("Failed to connect ftp server : ", err)
		return
	}

	data := bytes.NewBuffer(file)
	err = conn.Stor(filename, data)
	if err != nil {
		logrus.Error("Failed upload file to server : ", err)
		return
	}

	logrus.Info("File ", filename, " was upload to ", os.Getenv("ftp_addr"))
	return nil
}

// connectFTP connecting to ftp server
func connectFTP() (*ftp.ServerConn, error) {
	host := fmt.Sprintf("%s:%s", os.Getenv("ftp_addr"), os.Getenv("ftp_port"))
	username := fmt.Sprintf("%s", os.Getenv("ftp_username"))
	password := fmt.Sprintf("%s", os.Getenv("ftp_password"))

	conn, err := ftp.Dial(host)
	if err != nil {
		logrus.Error("Failed connect to ftp server : ", err)
		return nil, err

	}

	err = conn.Login(username, password)
	if err != nil {
		logrus.Error("Faile login to ftp server : ", err)
		return nil, err
	}

	return conn, err
}

//GetExtFile : getting extension file from base64Image Encode String
func GetExtFile(file string) string {
	slash := strings.Index(file, "/") + 1
	semicolon := strings.Index(file, ";")
	extImage := file[slash:semicolon]
	return extImage
}
