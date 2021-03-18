package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jabatan-struktural"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jabatan-strukural"
	"github.com/labstack/echo/v4"
)

func CreateJabatanStrukturalHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.JabatanStruktural)
	var jabatanStruktural *models.JabatanStruktural

	if err = c.Bind(&req); err != nil {
		return
	}

	repo := repository.NewJabatanStrukturalRepository()

	arg := repository.CreateJabatanStrukturalParams{
		Name:     req.Name,
		Fullname: req.FullName,
	}

	jabatanStruktural, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Success!"
	resp.Body = map[string]interface{}{
		"jabatan_struktural": jabatanStruktural,
	}

	return c.JSON(http.StatusCreated, resp)
}
