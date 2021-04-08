package action

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/login"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/login"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Login struct{}

func NewLoginHandler() *Login {
	return &Login{}
}

type LoginTokenClaim struct {
	*jwt.StandardClaims
	models.Login
}

func (lg *Login) LoginHandler(c echo.Context) (err error) {

	var resp helper.Response
	var login *models.Login

	err = mid.ValidationKey(c)
	if err != nil {
		return
	}

	if err = c.Bind(&login); err != nil {
		return
	}

	username := login.Username
	password := login.Password

	lowercaseUsername := strings.ToLower(username)

	repo := repository.NewLoginRepository()

	arg := repository.LoginParams{
		Username: lowercaseUsername,
		Password: password,
	}

	login, err = repo.Login(arg)
	if err != nil {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login gagal! Periksa kembali username dan password anda!"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &LoginTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		models.Login{
			ID:         login.ID,
			IDPengguna: login.IDPengguna,
			Username:   login.Username,
			CreatedAt:  login.CreatedAt,
			UpdatedAt:  login.UpdatedAt,
		},
	}

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
