package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jenis-surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jenis-surat"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type List struct{}

func NewListJenisSurat() *List {
	return &List{}
}

func (ls *List) ListJenisSuratHandler(c echo.Context) (err error) {

	var jenis []models.JenisSurat
	var resp helper.Response

	err = mid.ValidationKey(c)
	if err != nil {
		return
	}

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

	repo := repository.NewJenisSuratRepository()

	jenis, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"jenis": jenis,
	}

	return c.JSON(http.StatusOK, resp)
}
