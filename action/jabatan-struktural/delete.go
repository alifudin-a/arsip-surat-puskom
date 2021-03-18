package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jabatan-strukural"
	"github.com/labstack/echo/v4"
)

func DeleteJabatanStrukturalHandler(c echo.Context) (err error) {
	var resp helper.Response

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	repo := repository.NewJabatanStrukturalRepository()

	arg := repository.DeleteJabatanStrukturalParams{
		ID: int64(id),
	}

	exist, err := repo.IsExist(repository.IsExistJabatanStrukturalParams{ID: id})
	if err != nil {
		return
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
