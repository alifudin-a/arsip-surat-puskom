package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/login"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/login"
	"github.com/labstack/echo/v4"
)

type Login struct{}

func NewLoginHandler() *Login {
	return &Login{}
}

func (lg *Login) validate(req *models.Login, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (lg *Login) LoginHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.Login)
	var login *models.Login

	err = lg.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewLoginRepository()

	arg := repository.GetUsernameParams{
		Username: req.Username,
		Password: req.Password,
	}

	login, err = repo.GetUsername(arg)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Login gagal!")
	}

	// _ = login

	if req.Username != login.Username && req.Password != login.Password {
		resp.Code = http.StatusUnauthorized
		resp.Message = "Login gagal! Periksa kembali username dan password anda!"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	resp.Code = http.StatusOK
	resp.Message = "Login Berhasil!"

	return c.JSON(http.StatusOK, resp)
}
