package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Read struct{}

func NewReadSurat() *Read {
	return &Read{}
}

func (rd *Read) ReadSuratHandler(c echo.Context) (err error) {
	var resp helper.Response
	var surat *models.ReadSurat

	// err = mid.ValidationKey(c)
	// if err != nil {
	// 	return
	// }

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewSuratRepository()

	arg := repository.ReadSuratParams{
		ID: int64(id),
	}

	surat, err = repo.FindById(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"surat": surat,
	}

	return c.JSON(http.StatusOK, resp)
}
