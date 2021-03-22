package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
	"github.com/labstack/echo/v4"
)

type CreateV2 struct{}

func NewCreateSuratKeluarV2() *CreateV2 {
	return &CreateV2{}
}

func (cr *CreateV2) validate(req *models.CreateSuratKeluar, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *CreateV2) NewCreateSuratKeluarHandlerV2(c echo.Context) (err error) {
	var resp helper.Response
	var suratKeluar *models.CreateSuratKeluar
	var req = new(models.CreateSuratKeluar)

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratKeluarRepository()

	arg := builder.CreateSuratKeluarV2(req)

	suratKeluar, err = repo.CreateV2(arg)
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
