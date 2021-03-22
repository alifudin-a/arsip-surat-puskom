package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/penerima"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
	"github.com/labstack/echo/v4"
)

type Read struct{}

func NewReadPenerima() *Read {
	return &Read{}
}

func (rd *Read) ReadPenerimaHandler(c echo.Context) (err error) {
	var resp helper.Response

	var penerima = &models.Penerima{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewPenerimaRepository()

	arg := repository.ReadPenerimaParams{
		ID: int64(id),
	}

	penerima, err = repo.FindById(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"penerima": penerima,
	}

	return c.JSON(http.StatusOK, resp)
}
