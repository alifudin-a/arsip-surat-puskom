package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	"github.com/labstack/echo/v4"
)

type List struct {
}

func NewListSurat() *List {
	return &List{}
}

func (ls *List) ListSuratHandler(c echo.Context) (err error) {
	var resp helper.Response
	var surat []models.ListSurat

	repo := repository.NewSuratRepository()

	surat, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"surat": surat,
	}

	return c.JSON(http.StatusOK, resp)
}
