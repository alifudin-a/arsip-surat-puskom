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

func (cr *Create) validate(req *models.XCreateSurat, c echo.Context) (err error) {
	if err = c.Bind(req); err != nil {
		return
	}

	return c.Validate(req)
}

func (cr *Create) CreateSuratHandler(c echo.Context) (err error) {
	var resp helper.Response
	var req = new(models.XCreateSurat)

	err = cr.validate(req, c)
	if err != nil {
		return
	}

	arg := builder.CreateSurat(req)

	repo := repository.NewSuratRepository()

	var suratPenerima *models.XCreateSurat

	suratPenerima, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil menambahkan surat!"
	resp.Body = map[string]interface{}{
		"surat": suratPenerima,
	}

	return c.JSON(http.StatusOK, resp)
}
