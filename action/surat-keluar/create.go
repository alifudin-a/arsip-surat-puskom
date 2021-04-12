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

type Create struct {
}

func NewCreateSuratKeluar() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.CreateSuratKeluar, c echo.Context) (err error) {
	if err = c.Bind(req); err != nil {
		return
	}

	return c.Validate(req)
}

func (cr *Create) CreateSuratKeluarHandler(c echo.Context) (err error) {
	var resp helper.Response
	var req = new(models.CreateSuratKeluar)
	var suratKeluar *models.CreateSuratKeluar

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	err = cr.validate(req, c)
	if err != nil {
		return
	}

	arg := builder.CreateSuratKeluar(req)

	repo := repository.NewSuratKeluarRepository()

	suratKeluar, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil Membuat Surat Keluar!"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusCreated, resp)
}
