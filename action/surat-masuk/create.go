package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreateSuratMasuk() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.CreateSuratMasuk, c echo.Context) (err error) {
	if err = c.Bind(req); err != nil {
		return
	}

	return c.Validate(req)
}

func (cr *Create) CreateSuratMasukHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.CreateSuratMasuk)
	var suratMasuk *models.CreateSuratMasuk

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	err = cr.validate(req, c)
	if err != nil {
		return
	}

	repo := repository.NewSuratMasukRepository()

	arg := builder.CreateSuratMasuk(req)

	suratMasuk, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil menambahkan surat masuk!"
	resp.Body = map[string]interface{}{
		"surat_masuk": suratMasuk,
	}

	return c.JSON(http.StatusOK, resp)
}
