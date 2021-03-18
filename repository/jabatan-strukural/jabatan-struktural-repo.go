package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jabatan-struktural"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type JabatanStrukturalRepository interface {
	FindAll() ([]models.JabatanStruktural, error)
	FindById(arg ReadJabatanStrukturalParams) (*models.JabatanStruktural, error)
	Delete(arg DeleteJabatanStrukturalParams) (err error)
	IsExist(arg IsExistJabatanStrukturalParams) (bool, error)
	Create(arg CreateJabatanStrukturalParams) (*models.JabatanStruktural, error)
	Update(arg UpdateJabatanStrukturalParams) (*models.JabatanStruktural, error)
}

type repo struct{}

func NewJabatanStrukturalRepository() JabatanStrukturalRepository {
	return &repo{}
}

// ListJabatanStrukturalParams
// type ListJabatanStruktural struct{}

func (*repo) FindAll() ([]models.JabatanStruktural, error) {

	var jabatanStruktural []models.JabatanStruktural

	var db = database.OpenDB()

	err := db.Select(&jabatanStruktural, query.ListAllJabatanStruktural)
	if err != nil {
		return nil, err
	}

	return jabatanStruktural, nil
}

// ReadJabatanStrukturalParams .
type ReadJabatanStrukturalParams struct {
	ID int64
}

func (*repo) FindById(arg ReadJabatanStrukturalParams) (*models.JabatanStruktural, error) {
	var jabatanStruktural models.JabatanStruktural
	var db = database.OpenDB()

	err := db.Get(&jabatanStruktural, query.ReadJabatanStrukturalByID, arg.ID)
	if err != nil {
		return nil, err
	}

	return &jabatanStruktural, nil
}

// DeleteJabatanStrukturalParams .
type DeleteJabatanStrukturalParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteJabatanStrukturalParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteJabatanStruktural, arg.ID)
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

// IsExistJabatanStrukturalParams .
type IsExistJabatanStrukturalParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistJabatanStrukturalParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsExistJabatanStruktural, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

// CreateJabatanStrukturalParams .
type CreateJabatanStrukturalParams struct {
	Name     string
	Fullname string
}

func (*repo) Create(arg CreateJabatanStrukturalParams) (*models.JabatanStruktural, error) {
	var jabatanStruktural models.JabatanStruktural
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateJabatanStruktural, arg.Name, arg.Fullname).StructScan(&jabatanStruktural)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &jabatanStruktural, nil
}

// UpdateJabatanStrukturalParams .
type UpdateJabatanStrukturalParams struct {
	ID       int64
	Name     string
	Fullname string
}

func (*repo) Update(arg UpdateJabatanStrukturalParams) (*models.JabatanStruktural, error) {
	var jabatanStruktural models.JabatanStruktural
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateJabatanStruktural, arg.Name, arg.Fullname, arg.ID).StructScan(&jabatanStruktural)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &jabatanStruktural, nil
}
