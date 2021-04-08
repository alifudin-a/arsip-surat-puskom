package action

import (
	"net/http"
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jenis-surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jenis-surat"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreateJenisSurat() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.JenisSurat, c echo.Context) (err error) {

	if err = c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create) CreateJenisSuratHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.JenisSurat)
	var jenis *models.JenisSurat
	var t = time.Now()

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

	repo := repository.NewJenisSuratRepository()

	arg := repository.CreateJenisSuratParams{
		Name:      req.Name,
		CreatedAt: t.Format(helper.LayoutTime),
	}

	jenis, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil membuat jenis surat!"
	resp.Body = map[string]interface{}{
		"jenis": jenis,
	}

	return c.JSON(http.StatusCreated, resp)
}
