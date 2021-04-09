package action

import (
	"net/http"
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Update struct{}

func NewUpdateSuratMasuk() *Update {
	return &Update{}
}

func (up *Update) UpdateSuratMasukHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.SuratMasuk)
	var suratMasuk *models.SuratMasuk
	t := time.Now()

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	if err = c.Bind(req); err != nil {
		return
	}

	repo := repository.NewSuratMasukRepository()

	arg := repository.UpdateSuratMasukParams{
		Tanggal:    req.Tanggal,
		Nomor:      req.Nomor,
		IDPengirim: req.IDPengirim,
		Perihal:    req.Perihal,
		IDJenis:    *req.IDJenis,
		Keterangan: *req.Keterangan,
		UpdatedAt:  t.Format(helper.LayoutTime),
		ID:         req.ID,
	}

	suratMasuk, err = repo.Update(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil mengubah surat masuk!"
	resp.Body = map[string]interface{}{
		"surat_masuk": suratMasuk,
	}

	return c.JSON(http.StatusOK, resp)
}
