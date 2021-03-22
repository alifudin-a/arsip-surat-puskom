package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/penerima"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
	"github.com/labstack/echo/v4"
)

type List struct{}

func NewListPenerima() *List {
	return &List{}
}

func (ls *List) ListPenerimaHandler(c echo.Context) (err error) {

	var penerima []models.Penerima
	var resp helper.Response

	repo := repository.NewPenerimaRepository()

	penerima, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"penerima": penerima,
	}

	return c.JSON(http.StatusOK, resp)
}
