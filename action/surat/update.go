package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/builder"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Update struct{}

func NewUpdateSurat() *Update {
	return &Update{}
}

func (up *Update) validate(req *models.CreateSuratPenerima, c echo.Context) (err error) {
	if err := c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (up *Update) UpdateSuratHandler(c echo.Context) (err error) {
	var resp helper.Response
	var surat *models.CreateSuratPenerima
	var req = new(models.CreateSuratPenerima)

	// err = mid.ValidationKey(c)
	// if err != nil {
	// 	return
	// }

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

	err = up.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewSuratRepository()

	arg := builder.UpdateSurat(req)

	surat, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil mengubah surat !"
	resp.Body = map[string]interface{}{
		"surat_": surat,
	}

	return c.JSON(http.StatusOK, resp)
}
