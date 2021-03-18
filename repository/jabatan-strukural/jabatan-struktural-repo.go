package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/jabatan-struktural"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type JabatanStrukturalRepository interface {
	FindAll() ([]models.JabatanStruktural, error)
	FindById(arg ReadJabatanStrukturalParams) (*models.JabatanStruktural, error)
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
	ID int64 `json:"id"`
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
