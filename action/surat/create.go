package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreateSurat() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.Surat, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create) CreateSuratHandler(c echo.Context) (err error) {
	var resp helper.Response
	var surat *models.Surat
	var req = new(models.Surat)

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratRepository()

	arg := builder.CreateSurat(req)

	surat, err = repo.Create(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil menambahkan surat !"
	resp.Body = map[string]interface{}{
		"surat": surat,
	}

	return c.JSON(http.StatusOK, resp)
}
