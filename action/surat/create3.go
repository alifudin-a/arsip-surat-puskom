package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	"github.com/labstack/echo/v4"
)

type Create3 struct{}

func NewCreate3Surat() *Create3 {
	return &Create3{}
}

func (cr *Create3) validate(req *models.BuatSurat, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create3) CreateSurat3Handler(c echo.Context) (err error) {
	var resp helper.Response
	var surat *models.BuatSurat
	var req = new(models.BuatSurat)

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratRepository()

	arg := builder.CreateSurat3(req)

	surat, err = repo.BuatSurat(arg)
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
