package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
	"github.com/labstack/echo/v4"
)

type Read struct{}

func NewReadSuratKeluar() *Read {
	return &Read{}
}

func (rd *Read) ReadSuratKeluarHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratKeluar *models.ListSuratKeluar

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewSuratKeluarRepository()

	arg := repository.ReadSuratKeluarParams{
		ID: int64(id),
	}

	suratKeluar, err = repo.FindById(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusOK, resp)
}
