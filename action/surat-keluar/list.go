package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
	"github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type List struct{}

func NewListSuratKeluar() *List {
	return &List{}
}

func (ls *List) ListSuratKeluarByIDPengirim(c echo.Context) (err error) {

	var resp helper.Response
	var suratKeluar []models.ListSuratKeluar

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	repo := repository.NewSuratKeluarRepository()

	arg := repository.ListSuratKeluarByIDPengirimParams{
		IDPengguna: int64(id),
	}

	suratKeluar, err = repo.FindAllByIDPengguna(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil Menampilkan Surat Keluar!"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ls *List) ListSuratKeluarHandler(c echo.Context) (err error) {

	var resp helper.Response
	var suratKeluar []models.ListSuratKeluar
	var qparam = c.QueryParam("asc_offset")

	err = middleware.ValidationJWT(c)
	if err != nil {
		return
	}

	repo := repository.NewSuratKeluarRepository()

	arg := repository.ListSUratKeluarAscParams{
		Offset: 0,
	}

	if qparam != "" {
		suratKeluar, err = repo.FindAllAsc(arg, qparam)
	} else {
		suratKeluar, err = repo.FindAllDesc()
	}

	if err != nil {
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil Menampilkan Surat Keluar!"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusOK, resp)
}
