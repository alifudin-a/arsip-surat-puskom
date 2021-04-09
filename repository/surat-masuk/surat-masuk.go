package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type SuratMasukRepository interface {
	FindAllDesc() ([]models.ListSuratMasuk, error)
	FindAllAsc(arg ListSuratMasukParams, queryparam string) ([]models.ListSuratMasuk, error)
	FindByID(arg GetSuratMasukParams) (*models.ListSuratMasuk, error)
	Delete(arg DeleteSuratMasukParams) (err error)
	IsExist(arg IsSuratMasukExistParams) (bool, error)
}

type repo struct{}

func NewSuratMasukRepository() SuratMasukRepository {
	return &repo{}
}

func (*repo) FindAllDesc() ([]models.ListSuratMasuk, error) {

	var suratMasuk []models.ListSuratMasuk
	var db = database.OpenDB()

	err := db.Select(&suratMasuk, query.ListSuratMasukDesc)
	if err != nil {
		return nil, err
	}

	return suratMasuk, nil
}

type ListSuratMasukParams struct {
	Offset int64
}

func (*repo) FindAllAsc(arg ListSuratMasukParams, queryparam string) ([]models.ListSuratMasuk, error) {

	var suratMasuk []models.ListSuratMasuk
	var db = database.OpenDB()

	err := db.Select(&suratMasuk, query.ListSuratMasukAsc, arg.Offset)
	if err != nil {
		return nil, err
	}

	return suratMasuk, nil
}

type GetSuratMasukParams struct {
	ID int64
}

func (*repo) FindByID(arg GetSuratMasukParams) (*models.ListSuratMasuk, error) {

	var suratMasuk models.ListSuratMasuk
	var db = database.OpenDB()

	err := db.Get(&suratMasuk, query.GetSuratMasukByID, arg.ID)
	if err != nil {
		return nil, err
	}

	return &suratMasuk, nil
}

type DeleteSuratMasukParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteSuratMasukParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteSuratMasuk, arg.ID)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return nil
}

type IsSuratMasukExistParams struct {
	ID int64
}

func (*repo) IsExist(arg IsSuratMasukExistParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsSuratMasukExist, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}
