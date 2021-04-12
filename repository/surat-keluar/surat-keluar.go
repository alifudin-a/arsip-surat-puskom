package repository

import (
	"encoding/json"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
	"github.com/jmoiron/sqlx/types"
)

type SuratKeluarRepository interface {
	FindAllDesc() ([]models.ListSuratKeluar, error)
	FindAllAsc(arg ListSUratKeluarAscParams, queryparam string) ([]models.ListSuratKeluar, error)
	FindAllByIDPengguna(arg ListSuratKeluarByIDPengirimParams) ([]models.ListSuratKeluar, error)
	FindByID(arg ReadSuratKeluarParams) (*models.ReadSuratKeluar, error)
}

type repo struct{}

func NewSuratKeluarRepository() SuratKeluarRepository {
	return &repo{}
}

type ListSuratKeluarByIDPengirimParams struct {
	IDPengguna int64
}

func (*repo) FindAllByIDPengguna(arg ListSuratKeluarByIDPengirimParams) ([]models.ListSuratKeluar, error) {

	var suratKeluar []models.ListSuratKeluar
	var db = database.OpenDB()
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluarByIDPengirim, arg.IDPengguna)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

func (*repo) FindAllDesc() ([]models.ListSuratKeluar, error) {
	var suratKeluar []models.ListSuratKeluar
	var db = database.OpenDB()
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluar)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

type ListSUratKeluarAscParams struct {
	Offset int64
}

func (*repo) FindAllAsc(arg ListSUratKeluarAscParams, queryparam string) ([]models.ListSuratKeluar, error) {
	var suratKeluar []models.ListSuratKeluar
	var db = database.OpenDB()
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluarAsc, arg.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

type ReadSuratKeluarParams struct {
	ID int64
}

func (*repo) FindByID(arg ReadSuratKeluarParams) (*models.ReadSuratKeluar, error) {
	var suratKeluar models.ReadSuratKeluar
	var db = database.OpenDB()
	var jsonString types.JSONText

	row, err := db.Queryx(query.ReadSuratKeluar, arg.ID)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonString), &suratKeluar)
		if err != nil {
			return nil, err
		}
	}

	return &suratKeluar, nil
}
