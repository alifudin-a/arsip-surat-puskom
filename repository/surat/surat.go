package repository

import (
	"encoding/json"
	"fmt"
	"time"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
	"github.com/jmoiron/sqlx/types"
)

type SuratRepository interface {
	FindAll() ([]models.ListSurat, error)
	FindById(arg ReadSuratParams) (*models.ReadSurat, error)
	Delete(arg DeleteSuratParams) (err error)
	IsExist(arg IsExistSuratParams) (bool, error)
	Create(arg CreateSurat) (*models.XCreateSurat, error)
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

type CreateSurat struct {
	Surat    models.Surat
	Penerima []models.Penerima
}

func (r *repo) Create(arg CreateSurat) (*models.XCreateSurat, error) {
	var suratPenerima models.XCreateSurat
	var err error

	var surat *models.Surat
	surat, err = r.createSurat(&arg)
	if err != nil {
		return nil, err
	}

	suratPenerima.Surat = *surat
	arg.Surat.ID = surat.ID

	var penerima []models.Penerima
	penerima, err = r.createPenerima(&arg)
	if err != nil {
		return nil, err
	}
	suratPenerima.Penerima = penerima

	return &suratPenerima, nil
}

func (r *repo) createSurat(arg *CreateSurat) (*models.Surat, error) {
	var surat models.Surat
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.XCreateSurat,
		arg.Surat.Tanggal,
		arg.Surat.Nomor,
		arg.Surat.IDPengirim,
		arg.Surat.Perihal,
		arg.Surat.IDJenis,
		arg.Surat.Keterangan,
		arg.Surat.CreatedAt,
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

func (r *repo) createPenerima(arg *CreateSurat) ([]models.Penerima, error) {
	var suratPenerima []models.Penerima
	var db = database.OpenDB()

	q := query.XCreatePenerima
	t := time.Now()

	insertParams := []interface{}{}

	for i, v := range arg.Penerima {
		v.IDSurat = arg.Surat.ID
		v.CreatedAt2 = t.Format(helper.LayoutTime)

		var s models.Penerima

		s.IDSurat = v.IDSurat
		s.IDPengguna = v.IDPengguna
		s.CreatedAt2 = v.CreatedAt2

		p1 := i * 3
		q += fmt.Sprintf("($%d,$%d,$%d),", p1+1, p1+2, p1+3)
		insertParams = append(insertParams, v.IDSurat, v.IDPengguna, v.CreatedAt2)
		suratPenerima = append(suratPenerima, s)
	}

	q = q[:len(q)-1]

	tx := db.MustBegin()
	_, err := tx.Exec(q, insertParams...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return suratPenerima, err
}
