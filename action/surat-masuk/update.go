package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/labstack/echo/v4"
)

type Update struct{}

func NewUpdateSuratMasuk() *Update {
	return &Update{}
}

func (up *Update) validate(req *models.SuratMasuk, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (up *Update) NewUpdateSuratMasukHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratMasuk *models.SuratMasuk
	var req = new(models.SuratMasuk)

	err = up.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratMasukRepository()

	arg := builder.UpdateSuratMasuk(req)

	exist, err := repo.IsExist(repository.IsExistSuratMasukParams{
		ID: arg.ID,
	})
	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)

	}
	if err != nil {
		return err
	}

	suratMasuk, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil mengubah surat Masuk!"
	resp.Body = map[string]interface{}{
		"surat_masuk": suratMasuk,
	}

	return c.JSON(http.StatusOK, resp)
}
