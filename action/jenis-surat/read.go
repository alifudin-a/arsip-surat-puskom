package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jenis-surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jenis-surat"
	"github.com/labstack/echo/v4"
)

type Read struct{}

func NewReadJenisSurat() *Read {
	return &Read{}
}

func (rd *Read) ReadJenisSuratHandler(c echo.Context) (err error) {
	var resp helper.Response

	var jenis = &models.JenisSurat{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewJenisSuratRepository()

	arg := repository.ReadJenisSuratParams{
		ID: int64(id),
	}

	jenis, err = repo.FindById(arg)
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
