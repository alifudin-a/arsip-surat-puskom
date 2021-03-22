package action

import (
	"net/http"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/login"
	"github.com/labstack/echo/v4"
)

type LoginV2 struct{}

func NewLoginHandlerV2() *Login {
	return &Login{}
}
func (lg *Login) LoginHandlerV2(c echo.Context) (err error) {

	var resp helper.Response
	var login models.Login
	var db = database.OpenDB()

	if err = c.Bind(&login); err != nil {
		return err
	}

	username := login.Username
	password := login.Password

	err = db.Get(&login, "SELECT * FROM tbl_login WHERE username = $1", username)
	if err != nil {
		return err
	}

	// _ = login

	if username != login.Username && password != login.Password {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login gagal! Periksa kembali username dan password anda!"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	resp.Code = http.StatusOK
	resp.Message = "Login Berhasil!"
	resp.Body = map[string]interface{}{
		"login": login,
	}

	return c.JSON(http.StatusOK, resp)
}
