package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type List struct{}

func NewListSuratMasuk() *List {
	return &List{}
}

func (ls *List) ListSuratMasukHandler(c echo.Context) (err error) {

	var resp helper.Response
	var surat []models.ListSuratMasuk
	var qparam = c.QueryParam("asc_offset")

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	repo := repository.NewSuratMasukRepository()

	arg := repository.ListSuratMasukParams{
		Offset: 0,
	}

	if qparam != "" {
		surat, err = repo.FindAllAsc(arg, qparam)
	} else {
		surat, err = repo.FindAllDesc()
	}
	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil Menampilkan Surat Masuk!"
	resp.Body = map[string]interface{}{
		"surat_masuk": surat,
	}

	return c.JSON(http.StatusOK, resp)
}
