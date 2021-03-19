package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jenis-surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jenis-surat"
	"github.com/labstack/echo/v4"
)

func ListJenisSuratHandler(c echo.Context) (err error) {

	var jenis []models.JenisSurat
	var resp helper.Response

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
