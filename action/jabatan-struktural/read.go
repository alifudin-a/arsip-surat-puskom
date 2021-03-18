package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jabatan-struktural"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jabatan-strukural"
	"github.com/labstack/echo/v4"
)

func ReadJabatanStrukturalHandler(c echo.Context) (err error) {
	var resp helper.Response

	var jabatanStruktural = &models.JabatanStruktural{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewJabatanStrukturalRepository()

	arg := repository.ReadJabatanStrukturalParams{
		ID: int64(id),
	}

	jabatanStruktural, err = repo.FindById(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Success!"
	resp.Body = map[string]interface{}{
		"jabatan_struktural": jabatanStruktural,
	}

	return c.JSON(http.StatusOK, resp)
}
