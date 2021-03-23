package action

import (
	"net/http"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
	"github.com/labstack/echo/v4"
)

type ListJoin struct {
}

func NewListJoinSuratKeluar() *List {
	return &List{}
}

func (ls *List) ListJoinSuratKeluarHandler(c echo.Context) (err error) {
	var resp helper.Response

	var db = database.OpenDB()

	list, err := db.Queryx(query.SuratKeluarJoin)
	if err != nil {
		return err
	}

	joins := []models.SelectJoinSuratKeluar{}

	for list.Next() {
		var join models.SelectJoinSuratKeluar
		err = list.Scan(
			&join.ID,
			&join.Tanggal,
			&join.Nomor,
			&join.Pengirim,
			&join.Perihal,
			&join.Keterangan,
			&join.Penerima.ID,
			&join.Penerima.IDSurat,
			&join.Penerima.Penerima,
		)
		if err != nil {
			return err
		}

		joins = append(joins, join)
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"surat_keluar": joins,
	}

	return c.JSON(http.StatusOK, resp)
}
