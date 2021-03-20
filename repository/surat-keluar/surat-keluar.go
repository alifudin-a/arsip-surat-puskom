package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type SuratKeluarRepository interface {
	FindAll() ([]models.ListSuratKeluar, error)
	FindById(arg ReadSuratKeluarParams) (*models.ListSuratKeluar, error)
	Delete(arg DeleteSuratKeluarParams) (err error)
	IsExist(arg IsExistSuratKeluarParams) (bool, error)
	Create(arg CreateSuratKeluarParams) (*models.SuratKeluar, error)
	Update(arg UpdateSuratKeluarParams) (*models.SuratKeluar, error)
}

type repo struct{}

func NewSuratKeluarRepository() SuratKeluarRepository {
	return &repo{}
}

func (*repo) FindAll() ([]models.ListSuratKeluar, error) {

	var keluar []models.ListSuratKeluar

	var db = database.OpenDB()

	err := db.Select(&keluar, query.ListAllSuratKeluar)
	if err != nil {
		return nil, err
	}

	return keluar, nil
}

// ReadSuratKeluarParams .
type ReadSuratKeluarParams struct {
	ID int64
}

func (*repo) FindById(arg ReadSuratKeluarParams) (*models.ListSuratKeluar, error) {
	var keluar models.ListSuratKeluar
	var db = database.OpenDB()

	err := db.Get(&keluar, query.ReadSuratKeluarByID, arg.ID)
	if err != nil {
		return nil, err
	}

	return &keluar, nil
}

// DeleteSuratKeluarParams .
type DeleteSuratKeluarParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteSuratKeluarParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeletSuratKeluar, arg.ID)
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

// IsExistSuratKeluarParams .
type IsExistSuratKeluarParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistSuratKeluarParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsExistSuratKeluar, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

// CreateSuratKeluarParams .
type CreateSuratKeluarParams struct {
	Tanggal    string
	Nomor      string
	IDPenerima int64
	IDPengirim int64
	Perihal    string
	IDJenis    int64
	Keterangan string
	CreatedAt  string
}

func (*repo) Create(arg CreateSuratKeluarParams) (*models.SuratKeluar, error) {
	var keluar models.SuratKeluar
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateSuratKeluar,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPenerima,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.CreatedAt,
	).StructScan(&keluar)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &keluar, nil
}

// UpdateSuratKeluarParams .
type UpdateSuratKeluarParams struct {
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

func (*repo) Update(arg UpdateSuratKeluarParams) (*models.SuratKeluar, error) {
	var keluar models.SuratKeluar
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateSuratKeluar,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPenerima,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.UpdatedAt,
		arg.ID,
	).StructScan(&keluar)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &keluar, nil
}
