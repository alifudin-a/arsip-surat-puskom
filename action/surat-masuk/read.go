package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Read struct{}

func NewReadSuratMasuk() *Read {
	return &Read{}
}

func (rd *Read) ReadSuratMasukHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratMasuk *models.ListSuratMasuk

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewSuratMasukRepository()

	arg := repository.GetSuratMasukParams{
		ID: int64(id),
	}

	suratMasuk, err = repo.FindByID(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan surat masuk!"
	resp.Body = map[string]interface{}{
		"surat_masuk": suratMasuk,
	}

	return c.JSON(http.StatusOK, resp)
}
