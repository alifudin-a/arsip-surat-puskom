package action

import (
	"net/http"
	"os"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/login"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/login"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Login struct{}

func NewLoginHandler() *Login {
	return &Login{}
}

func (lg *Login) LoginHandler(c echo.Context) (err error) {

	var resp helper.Response
	var login *models.Login

	if err = c.Bind(&login); err != nil {
		return err
	}

	username := login.Username
	password := login.Password

	repo := repository.NewLoginRepository()

	arg := repository.LoginParams{
		Username: username,
		Password: password,
	}

	login, err = repo.Login(arg)
	if err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login gagal! Periksa kembali username dan password anda!"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	// Generate jwt token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = login.ID
	claims["id_pengguna"] = login.IDPengguna
	claims["username"] = login.Username
	claims["created_at"] = login.CreatedAt
	claims["updated_at"] = login.UpdatedAt

	tokenJwt, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Login Berhasil!"
	resp.Body = map[string]interface{}{
		"token": tokenJwt,
	}

	return c.JSON(http.StatusOK, resp)
}
