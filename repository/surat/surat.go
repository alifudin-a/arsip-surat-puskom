package repository

import (
	"encoding/json"
	"fmt"

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
	BuatSurat(arg BuatSuratParams) (*models.BuatSurat, error)
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

type BuatSuratParams struct {
	Surat    models.Surat3
	Penerima []models.Penerima
}

func (r *repo) BuatSurat(arg BuatSuratParams) (*models.BuatSurat, error) {

	var buatSurat models.BuatSurat

	var surat *models.Surat3
	surat, err := r.createSurat(&arg)
	if err != nil {
		return nil, err
	}

	buatSurat.Surat3 = *surat
	arg.Surat.ID = surat.ID

	var penerima []models.Penerima
	penerima, err = r.createPenerima(&arg)
	if err != nil {
		return nil, err
	}

	buatSurat.Penerima = penerima

	return &buatSurat, nil
}

func (*repo) createSurat(arg *BuatSuratParams) (*models.Surat3, error) {
	var surat models.Surat3
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateSurat,
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

func (*repo) createPenerima(arg *BuatSuratParams) ([]models.Penerima, error) {
	var penerima []models.Penerima
	var db = database.OpenDB()

	q := query.CreatePenerima

	insertParams := []interface{}{}

	for i, v := range arg.Penerima {
		v.IDSurat = arg.Surat.ID

		var p models.Penerima

		p.IDSurat = v.IDSurat
		p.IDPengguna = v.IDPengguna
		p.CreatedAt = v.CreatedAt

		p1 := i * 3
		q += fmt.Sprintf("($%d,$%d,$%d),", p1+1, p1+2, p1+3)
		insertParams = append(insertParams, v.IDSurat, v.IDPengguna, v.CreatedAt)
		penerima = append(penerima, p)
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

	return penerima, nil
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
