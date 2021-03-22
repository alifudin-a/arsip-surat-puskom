package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/penerima"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
	"github.com/labstack/echo/v4"
)

type Update struct{}

func NewUpdatePenerima() *Update {
	return &Update{}
}

func (up *Update) validate(req *models.Penerima, c echo.Context) (err error) {

	if err = c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (up *Update) UpdatePenerimaHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.Penerima)
	var penerima *models.Penerima

	err = up.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewPenerimaRepository()

	arg := builder.UpdatePenerima(req)

	exist, err := repo.IsExist(repository.IsExistPenerimaParams{
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

	penerima, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil mebgubah data!"
	resp.Body = map[string]interface{}{
		"penerima": penerima,
	}

	return c.JSON(http.StatusOK, resp)
}
