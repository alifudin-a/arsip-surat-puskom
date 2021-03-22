package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/penerima"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type PenerimaRepository interface {
	FindAll() ([]models.Penerima, error)
	FindById(arg ReadPenerimaParams) (*models.Penerima, error)
	Delete(arg DeletePenerimaParams) (err error)
	IsExist(arg IsExistPenerimaParams) (bool, error)
	Create(arg CreatePenerimaParams) (*models.Penerima, error)
	Update(arg UpdatePenerimaParams) (*models.Penerima, error)
}

type repo struct{}

func NewPenerimaRepository() PenerimaRepository {
	return &repo{}
}

func (*repo) FindAll() ([]models.Penerima, error) {

	var penerima []models.Penerima

	var db = database.OpenDB()

	err := db.Select(&penerima, query.ListAllPenerima)
	if err != nil {
		return nil, err
	}

	return penerima, nil
}

// ReadPenerimalParams .
type ReadPenerimaParams struct {
	ID int64
}

func (*repo) FindById(arg ReadPenerimaParams) (*models.Penerima, error) {
	var penerima models.Penerima
	var db = database.OpenDB()

	err := db.Get(&penerima, query.ReadPenerimaById, arg.ID)
	if err != nil {
		return nil, err
	}

	return &penerima, nil
}

// DeletePenerimaParams .
type DeletePenerimaParams struct {
	ID int64
}

func (*repo) Delete(arg DeletePenerimaParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeletePenerima, arg.ID)
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

// IsExistPenerimaParams .
type IsExistPenerimaParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistPenerimaParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsExistPenerima, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

// CreatePenerimaParams .
type CreatePenerimaParams struct {
	IDSurat    int64
	IDPengguna int64
	CreatedAt  string
}

func (*repo) Create(arg CreatePenerimaParams) (*models.Penerima, error) {
	var penerima models.Penerima
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreatePenerima, arg.IDSurat, arg.IDPengguna, arg.CreatedAt).StructScan(&penerima)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &penerima, nil
}

// UpdatePenerimaParams .
type UpdatePenerimaParams struct {
	ID         int64
	IDSurat    int64
	IDPengguna int64
	UpdatedAt  string
}

func (*repo) Update(arg UpdatePenerimaParams) (*models.Penerima, error) {
	var penerima models.Penerima
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdatePenerima, arg.IDSurat, arg.IDPengguna, arg.UpdatedAt, arg.ID).StructScan(&penerima)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &penerima, nil
}
