package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Update struct{}

func NewUpdateSuratKeluar() *Update {
	return &Update{}
}

func (up *Update) validate(req *models.CreateSuratKeluar, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (up *Update) UpdateSuratHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratKeluar *models.CreateSuratKeluar
	var req = new(models.CreateSuratKeluar)

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	err = up.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratKeluarRepository()

	arg := builder.UpdateSuratKeluar(req)

	suratKeluar, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil mengubah surat !"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusOK, resp)
}
