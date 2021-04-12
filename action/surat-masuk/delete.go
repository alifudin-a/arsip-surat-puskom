package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Delete struct{}

func NewDeleteSuratMasuk() *Delete {
	return &Delete{}
}

func (dl *Delete) DeleteSuratMasukHandler(c echo.Context) (err error) {

	var resp helper.Response

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewSuratMasukRepository()

	arg := repository.DeleteSuratMasukParams{
		ID: id,
	}

	arg2 := repository.DeletePenerimaSuratParams{
		IDSurat: id,
	}

	exist2, err := repo.IsPenerimaSuratExist(repository.IsPenerimaSuratExistParams{ID: id})
	if err != nil {
		return
	}

	exist, err := repo.IsSuratMasukExist(repository.IsSuratMasukExistParams{ID: id})
	if err != nil {
		return
	}

	if !exist2 {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data Penerima tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data Surat tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = repo.DeletePenerimaSurat(arg2)
	if err != nil {
		return
	}

	err = repo.Delete(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menghapus data!"

	return c.JSON(http.StatusOK, resp)
}
