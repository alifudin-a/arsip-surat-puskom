package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
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

	// err = mid.ValidationKey(c)
	// if err != nil {
	// 	return
	// }

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

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
