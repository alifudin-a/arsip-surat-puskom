package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreateSuratKeluar() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.SuratKeluar, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create) NewCreateSuratKeluarHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratKeluar *models.SuratKeluar
	var req = new(models.SuratKeluar)

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratKeluarRepository()

	arg := builder.CreateSuratKeluar(req)

	suratKeluar, err = repo.Create(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil menambahkan surat keluar!"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusOK, resp)
}
