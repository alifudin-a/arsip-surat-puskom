package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	"github.com/labstack/echo/v4"
)

type Delete struct{}

func NewDeleteSurat() *Delete {
	return &Delete{}
}

func (dl *Delete) DeleteSuratHandler(c echo.Context) (err error) {
	var resp helper.Response

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewSuratRepository()

	arg := repository.DeleteSuratParams{
		ID: id,
	}

	exist, err := repo.IsExist(repository.IsExistSuratParams{ID: id})
	if err != nil {
		return err
	}

	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)

	}

	err = repo.Delete(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menghapus data!"

	return c.JSON(http.StatusOK, resp)
}
