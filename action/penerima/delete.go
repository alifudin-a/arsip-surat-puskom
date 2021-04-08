package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Delete struct{}

func NewDeletePenerima() *Delete {
	return &Delete{}
}

func (dl *Delete) DeletePenerimaHandler(c echo.Context) (err error) {
	var resp helper.Response

	// err = mid.ValidationKey(c)
	// if err != nil {
	// 	return
	// }

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

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
