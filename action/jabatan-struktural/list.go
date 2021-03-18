package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jabatan-struktural"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jabatan-strukural"
	"github.com/labstack/echo/v4"
)

func ListJabatanStrukturalHandler(c echo.Context) (err error) {

	var jabatanStruktural []models.JabatanStruktural
	var resp helper.Response

	repo := repository.NewJabatanStrukturalRepository()

	jabatanStruktural, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Body = map[string]interface{}{
		"jabatan_struktural": jabatanStruktural,
	}

	return c.JSON(http.StatusOK, resp)
}
