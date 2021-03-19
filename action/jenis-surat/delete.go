package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jenis-surat"
	"github.com/labstack/echo/v4"
)

type Delete struct{}

func NewDeleteJenisSurat() *Delete {
	return &Delete{}
}

func (dl *Delete) DeleteJenisSuratHandler(c echo.Context) (err error) {
	var resp helper.Response

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	repo := repository.NewJenisSuratRepository()

	arg := repository.DeleteJenisSuratParams{
		ID: int64(id),
	}

	exist, err := repo.IsExist(repository.IsExistJenisSuratParams{ID: id})
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
