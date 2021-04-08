package action

import (
	"net/http"
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jenis-surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/jenis-surat"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Update struct{}

func NewUpdateJenisSurat() *Update {
	return &Update{}
}

func (up *Update) validate(req *models.JenisSurat, c echo.Context) (err error) {

	if err = c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (up *Update) UpdateUserHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.JenisSurat)
	var jenis *models.JenisSurat
	var t = time.Now()

	err = mid.ValidationKey(c)
	if err != nil {
		return
	}

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

	err = up.validate(req, c)
	if err != nil {
		return nil
	}

	repo := repository.NewJenisSuratRepository()

	arg := repository.UpdateJenisSuratParams{
		ID:        req.ID,
		Name:      req.Name,
		UpdatedAt: t.Format(helper.LayoutTime),
	}

	exist, err := repo.IsExist(repository.IsExistJenisSuratParams{
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

	jenis, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil mebgubah data!"
	resp.Body = map[string]interface{}{
		"jenis": jenis,
	}

	return c.JSON(http.StatusOK, resp)
}
