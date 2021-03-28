package repository

import (
	"encoding/json"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
	"github.com/jmoiron/sqlx/types"
)

type SuratRepository interface {
	FindAll() ([]models.ListSurat, error)
	FindById(arg ReadSuratParams) (*models.ReadSurat, error)
	Delete(arg DeleteSuratParams) (err error)
	IsExist(arg IsExistSuratParams) (bool, error)
	Create(arg CreateSuratParams) (*models.Surat, error)
	Create2(arg CreateSurat2Params) (*models.CreateSurat, error)
	Update(arg UpdateSuratParams) (*models.Surat, error)
}

type repo struct{}

func NewSuratRepository() SuratRepository {
	return &repo{}
}

func (*repo) FindAll() ([]models.ListSurat, error) {
	var surat []models.ListSurat
	var db = database.OpenDB()
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSurat)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSurat

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		surat = append(surat, s)
	}

	return surat, nil
}

// ReadSuratParams .
type ReadSuratParams struct {
	ID int64
}

func (*repo) FindById(arg ReadSuratParams) (*models.ReadSurat, error) {
	var surat models.ReadSurat
	var db = database.OpenDB()
	var jsonString types.JSONText

	row, err := db.Queryx(query.ReadSuratByID, arg.ID)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonString), &surat)
		if err != nil {
			return nil, err
		}
	}

	return &surat, nil
}

// DeleteSuratParams .
type DeleteSuratParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteSuratParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteSurat, arg.ID)
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

// IsExistSuratParams .
type IsExistSuratParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistSuratParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsExistSurat, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

type CreateSurat2Params struct {
	Tanggal    string
	Nomor      string
	IDPengirim int64
	Perihal    string
	IDJenis    int64
	Keterangan string
	CreatedAt  string
	IDSurat    interface{}
	IDPengguna int64
	CreatedAt2 string
}

func (*repo) Create2(arg CreateSurat2Params) (*models.CreateSurat, error) {
	var surat models.CreateSurat
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateSurat2,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.CreatedAt,
		arg.IDSurat,
		arg.IDPengguna,
		arg.CreatedAt2,
	).StructScan(&surat)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &surat, nil
}

// CreateSuratParams .
type CreateSuratParams struct {
	Tanggal    string
	Nomor      string
	IDPengirim int64
	Perihal    string
	IDJenis    int64
	Keterangan string
	CreatedAt  string
}

func (*repo) Create(arg CreateSuratParams) (*models.Surat, error) {
	var surat models.Surat
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateSurat,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.CreatedAt,
	).StructScan(&surat)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &surat, nil
}

// UpdateSuratParams .
type UpdateSuratParams struct {
	ID         int64
	Tanggal    string
	Nomor      string
	IDPengirim int64
	Perihal    string
	IDJenis    int64
	Keterangan string
	UpdatedAt  string
}

func (*repo) Update(arg UpdateSuratParams) (*models.Surat, error) {
	var surat models.Surat
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateSurat,
		arg.Tanggal,
		arg.Nomor,
		arg.IDPengirim,
		arg.Perihal,
		arg.IDJenis,
		arg.Keterangan,
		arg.UpdatedAt,
		arg.ID,
	).StructScan(&surat)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &surat, nil
}
