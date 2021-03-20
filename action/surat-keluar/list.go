package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
	"github.com/labstack/echo/v4"
)

type List struct {
}

func NewListSuratKeluar() *List {
	return &List{}
}

func (ls *List) ListSuratKeluarHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratKeluar []models.ListSuratKeluar

	repo := repository.NewSuratKeluarRepository()

	suratKeluar, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"surat_keluar": suratKeluar,
	}

	return c.JSON(http.StatusOK, resp)
}
