package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
	"github.com/labstack/echo/v4"
)

type List struct {
}

func NewListSuratMasuk() *List {
	return &List{}
}

func (ls *List) ListSuratMasukHandler(c echo.Context) (err error) {
	var resp helper.Response
	var suratMasuk []models.ListSuratMasuk

	repo := repository.NewSuratMasukRepository()

	suratMasuk, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"surat_masuk": suratMasuk,
	}

	return c.JSON(http.StatusOK, resp)
}
