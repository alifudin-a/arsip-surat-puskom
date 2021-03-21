package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type SuratMasukRepository interface {
	FindAll() ([]models.ListSuratMasuk, error)
	FindById(arg ReadSuratMasukParams) (*models.ListSuratMasuk, error)
	Delete(arg DeleteSuratMasukParams) (err error)
	IsExist(arg IsExistSuratMasukParams) (bool, error)
	Create(arg CreateSuratMasukParams) (*models.SuratMasuk, error)
	Update(arg UpdateSuratMasukParams) (*models.SuratMasuk, error)
}

type repo struct{}

func NewSuratMasukRepository() SuratMasukRepository {
	return &repo{}
}

func (*repo) FindAll() ([]models.ListSuratMasuk, error) {

	var masuk []models.ListSuratMasuk

	var db = database.OpenDB()

	err := db.Select(&masuk, query.ListAllSuratMasuk)
	if err != nil {
		return nil, err
	}

	return masuk, nil
}

// ReadSuratMasukParams .
type ReadSuratMasukParams struct {
	ID int64
}

func (*repo) FindById(arg ReadSuratMasukParams) (*models.ListSuratMasuk, error) {
	var masuk models.ListSuratMasuk
	var db = database.OpenDB()

	err := db.Get(&masuk, query.ReadSuratMasukByID, arg.ID)
	if err != nil {
		return nil, err
	}

	return &masuk, nil
}

// DeleteSuratMasukParams .
type DeleteSuratMasukParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteSuratMasukParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeletSuratMasuk, arg.ID)
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

// IsExistSuratMasukParams .
type IsExistSuratMasukParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistSuratMasukParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsExistSuratMasuk, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

// CreateSuratMasukParams .
type CreateSuratMasukParams struct {
	Tanggal    string
	Nomor      string
	IDPenerima int64
	IDPengirim int64
	Perihal    string
	IDJenis    int64
	Keterangan string
	CreatedAt  string
}

func (*repo) Create(arg CreateSuratMasukParams) (*models.SuratMasuk, error) {
	var masuk models.SuratMasuk
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateSuratMasuk,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPenerima,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.CreatedAt,
	).StructScan(&masuk)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &masuk, nil
}

// UpdateSuratMasukParams .
type UpdateSuratMasukParams struct {
	ID         int64
	Tanggal    string
	Nomor      string
	IDPenerima int64
	IDPengirim int64
	Perihal    string
	IDJenis    int64
	Keterangan string
	UpdatedAt  string
}

func (*repo) Update(arg UpdateSuratMasukParams) (*models.SuratMasuk, error) {
	var masuk models.SuratMasuk
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateSuratMasuk,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPenerima,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.UpdatedAt,
		arg.ID,
	).StructScan(&masuk)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &masuk, nil
}
