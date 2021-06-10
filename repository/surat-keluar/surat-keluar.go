package repository

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	rd "github.com/alifudin-a/arsip-surat-puskom/database/redis"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lib/pq"
	errs "github.com/pkg/errors"
)

type SuratKeluarRepository interface {
	FindAllDesc() ([]models.ListSuratKeluar, error)
	FindAllAsc(arg ListSUratKeluarAscParams, queryparam string) ([]models.ListSuratKeluar, error)
	FindAllByIDPengirim(arg ListSuratKeluarByIDPengirimParams) ([]models.ListSuratKeluar, error)
	FindAllByIDPengirimAsc(arg ListSuratKeluarByIDPengirimAscParams, queryparam string) ([]models.ListSuratKeluar, error)
	FindByID(arg ReadSuratKeluarParams) (*models.ReadSuratKeluar, error)
	Delete(arg DeleteSuratKeluarParams) (err error)
	DeletePenerimaSuratKeluar(arg DeletePenerimaSuratKeluarParams) (err error)
	IsSuratMasukExist(arg IsSuratKeluarExistParams) (bool, error)
	IsPenerimaSuratExist(arg IsPenerimaSuratKeluarExistParams) (bool, error)
	Create(arg CreateSuratKeluarParams) (*models.CreateSuratKeluar, error)
	Update(arg UpdateSuratKeluarParams) (*models.CreateSuratKeluar, error)
	FindByIDandIDPengguna(arg FindByIDandIDPenggunaParams) (*models.ReadSuratKeluar, error)
}

type repo struct{}

func NewSuratKeluarRepository() SuratKeluarRepository {
	return &repo{}
}

type FindByIDandIDPenggunaParams struct {
	IDPengguna int64
	ID         int64
}

func (*repo) FindByIDandIDPengguna(arg FindByIDandIDPenggunaParams) (*models.ReadSuratKeluar, error) {

	var suratKeluar models.ReadSuratKeluar
	var db = database.DB
	var jsonString types.JSONText

	row, err := db.Queryx(query.ReadSuratKeluarByIDPenggunaByIDSuratKeluar, arg.IDPengguna, arg.ID)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonString), &suratKeluar)
		if err != nil {
			return nil, err
		}
	}

	return &suratKeluar, nil
}

type ListSuratKeluarByIDPengirimParams struct {
	IDPengguna int64
}

func (*repo) FindAllByIDPengirim(arg ListSuratKeluarByIDPengirimParams) ([]models.ListSuratKeluar, error) {

	var suratKeluar []models.ListSuratKeluar
	var db = database.DB
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluarByIDPengirim, arg.IDPengguna)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

type ListSuratKeluarByIDPengirimAscParams struct {
	IDPengguna int64
	Offset     int64
}

func (*repo) FindAllByIDPengirimAsc(arg ListSuratKeluarByIDPengirimAscParams, queryparam string) ([]models.ListSuratKeluar, error) {

	var suratKeluar []models.ListSuratKeluar
	var db = database.DB
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluarByIDPengirimAsc, arg.IDPengguna, arg.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

func (*repo) FindAllDesc() ([]models.ListSuratKeluar, error) {
	var suratKeluar []models.ListSuratKeluar
	var db = database.DB
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluar)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

type ListSUratKeluarAscParams struct {
	Offset int64
}

func (*repo) FindAllAsc(arg ListSUratKeluarAscParams, queryparam string) ([]models.ListSuratKeluar, error) {
	var suratKeluar []models.ListSuratKeluar
	var db = database.DB
	var jsonString types.JSONText

	rows, err := db.Queryx(query.ListSuratKeluarAsc, arg.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		var s models.ListSuratKeluar

		err = json.Unmarshal([]byte(jsonString), &s)
		if err != nil {
			return nil, err
		}

		suratKeluar = append(suratKeluar, s)
	}

	return suratKeluar, nil
}

type ReadSuratKeluarParams struct {
	ID int64
}

func (*repo) FindByID(arg ReadSuratKeluarParams) (*models.ReadSuratKeluar, error) {
	var suratKeluar models.ReadSuratKeluar
	var db = database.DB
	var jsonString types.JSONText

	row, err := db.Queryx(query.ReadSuratKeluar, arg.ID)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(&jsonString)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(jsonString), &suratKeluar)
		if err != nil {
			return nil, err
		}
	}

	return &suratKeluar, nil
}

type DeleteSuratKeluarParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteSuratKeluarParams) (err error) {
	var db = database.DB

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteSuratKeluar, arg.ID)
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

type DeletePenerimaSuratKeluarParams struct {
	IDSurat int64
}

func (*repo) DeletePenerimaSuratKeluar(arg DeletePenerimaSuratKeluarParams) (err error) {
	var db = database.DB

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeletePenerimaSuratKeluar, arg.IDSurat)
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

type IsSuratKeluarExistParams struct {
	ID int64
}

func (*repo) IsSuratMasukExist(arg IsSuratKeluarExistParams) (bool, error) {
	var db = database.DB
	var total int

	err := db.Get(&total, query.IsSuratKeluarExist, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

type IsPenerimaSuratKeluarExistParams struct {
	ID int64
}

func (*repo) IsPenerimaSuratExist(arg IsPenerimaSuratKeluarExistParams) (bool, error) {
	var db = database.DB
	var total int

	err := db.Get(&total, query.IsPenerimaSuratKeluarExist, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

type UpdateSuratKeluarParams struct {
	SuratKeluar         models.SuratKeluar
	PenerimaSuratKeluar []models.PenerimaSuratKeluar
}

func (r *repo) Update(arg UpdateSuratKeluarParams) (*models.CreateSuratKeluar, error) {

	var suratKeluar models.CreateSuratKeluar
	var err error

	var s1 *models.SuratKeluar
	s1, err = r.updateSurat(&arg)
	if err != nil {
		return nil, err
	}

	suratKeluar.SuratKeluar = *s1

	var s2 []models.PenerimaSuratKeluar
	s2, err = r.updatePenerima(&arg)
	if err != nil {
		return nil, err
	}

	suratKeluar.PenerimaSuratKeluar = s2

	return &suratKeluar, nil
}

func (*repo) updateSurat(arg *UpdateSuratKeluarParams) (*models.SuratKeluar, error) {
	var surat models.SuratKeluar
	var db = database.DB
	var err error

	var byteUpload []byte
	var filename string
	var fullpath string

	uploadPayload := arg.SuratKeluar.Upload
	if uploadPayload != "" {

		str := strings.SplitAfter(string(*&uploadPayload), ",")
		extFile := helper.GetExtFile(str[0])

		if extFile == "png" {
			byteUpload, err = base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(str[1]) //WithPadding(base64.NoPadding)
			if err != nil {
				return nil, err
			}
		} else {
			byteUpload, err = base64.StdEncoding.DecodeString(str[1])
			if err != nil {
				return nil, err
			}
		}

		filename = "surat_keluar_" + time.Now().Format(helper.LayoutTime3) + "." + extFile
		fullpath = "http://" + os.Getenv("ftp_addr") + ":" + os.Getenv("ftp_port_image") + "/" + filename

		arg.SuratKeluar.Upload = helper.NullString(*&fullpath)

		err = helper.Upload(byteUpload, filename)
		if err != nil {
			return nil, err
		}
	}

	tx := db.MustBegin()
	err = tx.QueryRowx(query.UpdateSuratKeluar,
		arg.SuratKeluar.Tanggal,
		arg.SuratKeluar.Nomor,
		arg.SuratKeluar.IDPengirim,
		arg.SuratKeluar.Perihal,
		arg.SuratKeluar.IDJenis,
		arg.SuratKeluar.Keterangan,
		arg.SuratKeluar.UpdatedAt,
		arg.SuratKeluar.Upload,
		arg.SuratKeluar.ID,
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

func (*repo) updatePenerima(arg *UpdateSuratKeluarParams) ([]models.PenerimaSuratKeluar, error) {
	var penerima []models.PenerimaSuratKeluar
	var db = database.DB
	var err error

	_, err = db.Exec("DELETE FROM tbl_penerima WHERE id_surat = $1", arg.SuratKeluar.ID)
	if err != nil {
		return penerima, errs.Wrap(err, "Gagal mengubah penerima!")
	}

	var surat models.SuratKeluar
	err = db.Get(&surat, "SELECT created_at FROM tbl_surat WHERE id = $1", arg.SuratKeluar.ID)
	if err != nil {
		return nil, err
	}

	q := query.UpdatePenerimaSuratKeluar

	updateParams := []interface{}{}

	for i, v := range arg.PenerimaSuratKeluar {
		v.IDSurat = arg.SuratKeluar.ID
		v.CreatedAt2 = surat.CreatedAt
		v.UpdatedAt2 = arg.SuratKeluar.UpdatedAt

		var s models.PenerimaSuratKeluar

		s.IDSurat = v.IDSurat
		s.IDPengguna = v.IDPengguna
		s.CreatedAt2 = v.CreatedAt2
		s.UpdatedAt2 = v.UpdatedAt2

		p1 := i * 4
		q += fmt.Sprintf("($%d,unnest(array[$%d::smallint[]]),$%d,$%d),", p1+1, p1+2, p1+3, p1+4)
		updateParams = append(updateParams, v.IDSurat, pq.Int64Array(v.IDPengguna), v.CreatedAt2, v.UpdatedAt2)
		penerima = append(penerima, s)
	}

	q = q[:len(q)-1]

	tx := db.MustBegin()
	_, err = tx.Exec(q, updateParams...)
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

type CreateSuratKeluarParams struct {
	SuratKeluar         models.SuratKeluar
	PenerimaSuratKeluar []models.PenerimaSuratKeluar
}

func (r *repo) Create(arg CreateSuratKeluarParams) (*models.CreateSuratKeluar, error) {

	var suratKeluar models.CreateSuratKeluar
	var err error

	var surat *models.SuratKeluar
	surat, err = r.createSurat(&arg)
	if err != nil {
		return nil, err
	}

	suratKeluar.SuratKeluar = *surat
	arg.SuratKeluar.ID = surat.ID

	var penerima []models.PenerimaSuratKeluar
	penerima, err = r.createPenerima(&arg)
	if err != nil {
		return nil, err
	}
	suratKeluar.PenerimaSuratKeluar = penerima

	return &suratKeluar, nil
}

func (r *repo) createSurat(arg *CreateSuratKeluarParams) (*models.SuratKeluar, error) {
	var surat models.SuratKeluar
	var db = database.DB
	var err error

	var byteUpload []byte

	uploadPayload := arg.SuratKeluar.Upload

	if uploadPayload != "" {

		err = rd.RdSet(arg.SuratKeluar.Perihal, string(uploadPayload), 0)
		if err != nil {
			return nil, err
		}

		str := strings.SplitAfter(string(*&uploadPayload), ",")
		extFile := helper.GetExtFile(str[0])

		if extFile == "png" {
			byteUpload, err = base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(str[1]) //WithPadding(base64.NoPadding)
			if err != nil {
				return nil, err
			}
		} else {
			byteUpload, err = base64.StdEncoding.DecodeString(str[1])
			if err != nil {
				return nil, err
			}
		}

		filename := "UNTUK_DIHAPUS_surat_keluar_" + time.Now().Format(helper.LayoutTime3) + "." + extFile

		fullpath := "http://" + os.Getenv("ftp_addr") + ":" + os.Getenv("ftp_port_image") + "/" + filename

		arg.SuratKeluar.Upload = helper.NullString(*&fullpath)

		err = helper.Upload(byteUpload, filename)
		if err != nil {
			return nil, err
		}
	}

	tx := db.MustBegin()
	err = tx.QueryRowx(query.CreateSuratKeluar,
		arg.SuratKeluar.Tanggal,
		arg.SuratKeluar.Nomor,
		arg.SuratKeluar.IDPengirim,
		arg.SuratKeluar.Perihal,
		arg.SuratKeluar.IDJenis,
		arg.SuratKeluar.Keterangan,
		arg.SuratKeluar.CreatedAt,
		arg.SuratKeluar.Upload,
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

func (r *repo) createPenerima(arg *CreateSuratKeluarParams) ([]models.PenerimaSuratKeluar, error) {
	var penerima []models.PenerimaSuratKeluar
	var db = database.DB

	q := query.CreatePenerimaSuratKeluar
	t := time.Now()

	insertParams := []interface{}{}

	for i, v := range arg.PenerimaSuratKeluar {
		v.IDSurat = arg.SuratKeluar.ID
		v.CreatedAt2 = helper.NullString(t.Format(helper.LayoutTime))

		var s models.PenerimaSuratKeluar

		s.IDSurat = v.IDSurat
		s.IDPengguna = v.IDPengguna
		s.CreatedAt2 = v.CreatedAt2

		p1 := i * 3
		q += fmt.Sprintf("($%d,unnest(array[$%d::smallint[]]),$%d),", p1+1, p1+2, p1+3)
		insertParams = append(insertParams, v.IDSurat, pq.Int64Array(v.IDPengguna), v.CreatedAt2)
		penerima = append(penerima, s)
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

	return penerima, err
}
