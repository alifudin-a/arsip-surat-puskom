package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreateSuratMasuk() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.SuratMasuk, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create) NewCreateSuratMasukHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratMasuk *models.SuratMasuk
	var req = new(models.SuratMasuk)

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratMasukRepository()

	arg := builder.CreateSuratMasuk(req)

	suratMasuk, err = repo.Create(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil menambahkan surat Masuk!"
	resp.Body = map[string]interface{}{
		"surat_masuk": suratMasuk,
	}

	return c.JSON(http.StatusOK, resp)
}
