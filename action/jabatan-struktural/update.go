package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jabatan-struktural"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jabatan-strukural"
	"github.com/labstack/echo/v4"
)

func UpdateJabatanStrukturalHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.JabatanStruktural)
	var JabatanStruktural *models.JabatanStruktural

	if err = c.Bind(&req); err != nil {
		return
	}

	repo := repository.NewJabatanStrukturalRepository()

	arg := repository.UpdateJabatanStrukturalParams{
		ID:       req.ID,
		Name:     req.Name,
		Fullname: req.FullName,
	}

	exist, err := repo.IsExist(repository.IsExistJabatanStrukturalParams{
		ID: int64(arg.ID),
	})
	if err != nil {
		return err
	}

	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)

	}

	JabatanStruktural, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Success!"
	resp.Body = map[string]interface{}{
		"jabatan_struktural": JabatanStruktural,
	}

	return c.JSON(http.StatusOK, resp)
}
