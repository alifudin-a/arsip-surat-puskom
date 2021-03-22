package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
	"github.com/labstack/echo/v4"
)

type Delete struct{}

func NewDeletePenerima() *Delete {
	return &Delete{}
}

func (dl *Delete) DeletePenerimaHandler(c echo.Context) (err error) {
	var resp helper.Response

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	repo := repository.NewPenerimaRepository()

	arg := repository.DeletePenerimaParams{
		ID: int64(id),
	}

	exist, err := repo.IsExist(repository.IsExistPenerimaParams{ID: id})
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
