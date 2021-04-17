package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jenis-surat"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type JenisSuratRepository interface {
	FindAll() ([]models.JenisSurat, error)
	FindById(arg ReadJenisSuratParams) (*models.JenisSurat, error)
	Delete(arg DeleteJenisSuratParams) (err error)
	IsExist(arg IsExistJenisSuratParams) (bool, error)
	Create(arg CreateJenisSuratParams) (*models.JenisSurat, error)
	Update(arg UpdateJenisSuratParams) (*models.JenisSurat, error)
}

type repo struct{}

func NewJenisSuratRepository() JenisSuratRepository {
	return &repo{}
}

// ListUserParams
// type ListUser struct{}

func (*repo) FindAll() ([]models.JenisSurat, error) {

	var jenis []models.JenisSurat

	var db = database.DB

	err := db.Select(&jenis, query.ListAllJenisSurat)
	if err != nil {
		return nil, err
	}

	return jenis, nil
}

// ReadUserlParams .
type ReadJenisSuratParams struct {
	ID int64
}

func (*repo) FindById(arg ReadJenisSuratParams) (*models.JenisSurat, error) {
	var jenis models.JenisSurat
	var db = database.DB

	err := db.Get(&jenis, query.ReadJenisSuratById, arg.ID)
	if err != nil {
		return nil, err
	}

	return &jenis, nil
}

// DeleteJenisSuratParams .
type DeleteJenisSuratParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteJenisSuratParams) (err error) {
	var db = database.DB

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteJenisSurat, arg.ID)
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

// IsExistUserParams .
type IsExistJenisSuratParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistJenisSuratParams) (bool, error) {
	var db = database.DB
	var total int

	err := db.Get(&total, query.IsExistJenisSurat, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

// CreateUserParams .
type CreateJenisSuratParams struct {
	Name      string
	CreatedAt string
}

func (*repo) Create(arg CreateJenisSuratParams) (*models.JenisSurat, error) {
	var jenis models.JenisSurat
	var db = database.DB

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateJenisSurat, arg.Name, arg.CreatedAt).StructScan(&jenis)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &jenis, nil
}

// UpdateUserParams .
type UpdateJenisSuratParams struct {
	Name      string
	UpdatedAt string
	ID        int64
}

func (*repo) Update(arg UpdateJenisSuratParams) (*models.JenisSurat, error) {
	var jenis models.JenisSurat
	var db = database.DB

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateJenisSurat, arg.Name, arg.UpdatedAt, arg.ID).StructScan(&jenis)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &jenis, nil
}
