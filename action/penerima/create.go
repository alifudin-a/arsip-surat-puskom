package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/penerima"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreatePenerima() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.Penerima, c echo.Context) (err error) {

	if err = c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create) CreatePenerimaHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.Penerima)
	var penerima *models.Penerima

	err = mid.ValidationKey(c)
	if err != nil {
		return
	}

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewPenerimaRepository()

	arg := builder.CreatePenerima(req)

	penerima, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil membuat Penerima!"
	resp.Body = map[string]interface{}{
		"penerima": penerima,
	}

	return c.JSON(http.StatusCreated, resp)
}
