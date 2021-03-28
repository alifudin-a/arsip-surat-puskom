package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	"github.com/labstack/echo/v4"
)

type Create2 struct{}

func NewCreate2Surat() *Create2 {
	return &Create2{}
}

func (cr *Create2) validate(req *models.CreateSurat, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create2) CreateSurat2Handler(c echo.Context) (err error) {
	var resp helper.Response
	var surat *models.CreateSurat
	var req = new(models.CreateSurat)

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratRepository()

	arg := builder.CreateSurat2(req)

	surat, err = repo.Create2(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil menambahkan surat!"
	resp.Body = map[string]interface{}{
		"surat": surat,
	}

	return c.JSON(http.StatusOK, resp)
}
