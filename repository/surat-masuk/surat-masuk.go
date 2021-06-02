package repository

import (
	"encoding/base64"
	"os"
	"strings"
	"time"

	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
	_ "github.com/joho/godotenv/autoload"
)

type SuratMasukRepository interface {
	FindAllDesc() ([]models.ListSuratMasuk, error)
	FindAllAsc(arg ListSuratMasukParams, queryparam string) ([]models.ListSuratMasuk, error)
	FindByID(arg GetSuratMasukParams) (*models.ListSuratMasuk, error)
	Delete(arg DeleteSuratMasukParams) (err error)
	DeletePenerimaSurat(arg DeletePenerimaSuratParams) (err error)
	IsSuratMasukExist(arg IsSuratMasukExistParams) (bool, error)
	IsPenerimaSuratExist(arg IsPenerimaSuratExistParams) (bool, error)
	FindAllByIDPengguna(arg ListSuratMasukByIDPenerimaParams) ([]models.ListSuratMasuk, error)
	FindAllByIDPenggunaAsc(arg ListSuratMasukByIDPenerimaAscParams, queryparam string) ([]models.ListSuratMasuk, error)
	Create(arg CreateSuratMasukParams) (*models.CreateSuratMasuk, error)
	Update(arg UpdateSuratMasukParams) (*models.CreateSuratMasuk, error)
}

type repo struct{}

func NewSuratMasukRepository() SuratMasukRepository {
	return &repo{}
}

// type FindByIDPenggunaAndIDParams struct {
// 	IDPengguna int64
// 	ID         int64
// }

// func (*repo) FindByIDPenggunaAndID(arg FindByIDPenggunaAndIDParams) ()

func (*repo) FindAllDesc() ([]models.ListSuratMasuk, error) {

	var suratMasuk []models.ListSuratMasuk
	var db = database.DB

	err := db.Select(&suratMasuk, query.ListSuratMasukDesc)
	if err != nil {
		return nil, err
	}

	return suratMasuk, nil
}

type ListSuratMasukParams struct {
	Offset int64
}

func (*repo) FindAllAsc(arg ListSuratMasukParams, queryparam string) ([]models.ListSuratMasuk, error) {

	var suratMasuk []models.ListSuratMasuk
	var db = database.DB

	err := db.Select(&suratMasuk, query.ListSuratMasukAsc, arg.Offset)
	if err != nil {
		return nil, err
	}

	return suratMasuk, nil
}

type GetSuratMasukParams struct {
	ID int64
}

func (*repo) FindByID(arg GetSuratMasukParams) (*models.ListSuratMasuk, error) {

	var suratMasuk models.ListSuratMasuk
	var db = database.DB

	err := db.Get(&suratMasuk, query.GetSuratMasukByID, arg.ID)
	if err != nil {
		return nil, err
	}

	return &suratMasuk, nil
}

type DeleteSuratMasukParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteSuratMasukParams) (err error) {
	var db = database.DB

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteSuratMasuk, arg.ID)
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

type DeletePenerimaSuratParams struct {
	IDSurat int64
}

func (*repo) DeletePenerimaSurat(arg DeletePenerimaSuratParams) (err error) {
	var db = database.DB

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeletePenerimaSurat, arg.IDSurat)
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

type IsSuratMasukExistParams struct {
	ID int64
}

func (*repo) IsSuratMasukExist(arg IsSuratMasukExistParams) (bool, error) {
	var db = database.DB
	var total int

	err := db.Get(&total, query.IsSuratMasukExist, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

type IsPenerimaSuratExistParams struct {
	ID int64
}

func (*repo) IsPenerimaSuratExist(arg IsPenerimaSuratExistParams) (bool, error) {
	var db = database.DB
	var total int

	err := db.Get(&total, query.IsPenerimaSuratExist, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

type ListSuratMasukByIDPenerimaParams struct {
	IDPengguna int64
}

func (*repo) FindAllByIDPengguna(arg ListSuratMasukByIDPenerimaParams) ([]models.ListSuratMasuk, error) {

	var suratMasuk []models.ListSuratMasuk
	var db = database.DB

	err := db.Select(&suratMasuk, query.ListSuratMasukByIDPenerima, arg.IDPengguna)
	if err != nil {
		return nil, err
	}

	return suratMasuk, nil
}

type ListSuratMasukByIDPenerimaAscParams struct {
	IDPengguna int64
	Offset     int64
}

func (*repo) FindAllByIDPenggunaAsc(arg ListSuratMasukByIDPenerimaAscParams, queryparam string) ([]models.ListSuratMasuk, error) {

	var suratMasuk []models.ListSuratMasuk
	var db = database.DB

	err := db.Select(&suratMasuk, query.ListSuratMasukByIDPenerimaAsc, arg.IDPengguna, arg.Offset)
	if err != nil {
		return nil, err
	}

	return suratMasuk, nil
}

type CreateSuratMasukParams struct {
	SuratMasuk models.SuratMasuk
	Penerima   models.Penerima
}

func (r *repo) Create(arg CreateSuratMasukParams) (*models.CreateSuratMasuk, error) {

	var suratMasuk models.CreateSuratMasuk
	var err error

	var s1 *models.SuratMasuk
	s1, err = r.createSurat(&arg)
	if err != nil {
		return nil, err
	}

	suratMasuk.SuratMasuk = *s1

	var s2 *models.Penerima
	arg.Penerima.IDSurat = s1.ID
	s2, err = r.createPenerima(&arg)
	if err != nil {
		return nil, err
	}

	suratMasuk.Penerima = *s2

	return &suratMasuk, nil
}

func (*repo) createSurat(arg *CreateSuratMasukParams) (*models.SuratMasuk, error) {
	var suratMasuk models.SuratMasuk
	var db = database.DB
	var err error

	uploadPayload := arg.SuratMasuk.Upload
	str := strings.SplitAfter(*uploadPayload, ",")
	extFile := helper.GetExtFile(str[0])

	var byteUpload []byte

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

	filename := "surat_masuk_" + time.Now().Format(helper.LayoutTime3) + "." + extFile
	fullpath := "http://" + os.Getenv("ftp_addr") + ":" + os.Getenv("ftp_port_image") + "/" + filename

	arg.SuratMasuk.Upload = &fullpath

	tx := db.MustBegin()
	err = tx.QueryRowx(query.CreateSuratMasuk,
		arg.SuratMasuk.Tanggal,
		arg.SuratMasuk.Nomor,
		arg.SuratMasuk.IDPengirim,
		arg.SuratMasuk.Perihal,
		arg.SuratMasuk.IDJenis,
		arg.SuratMasuk.Keterangan,
		arg.SuratMasuk.CreatedAt,
		arg.SuratMasuk.Upload,
	).StructScan(&suratMasuk)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	err = helper.Upload(byteUpload, filename)
	if err != nil {
		return nil, err
	}

	return &suratMasuk, nil
}

func (*repo) createPenerima(arg *CreateSuratMasukParams) (*models.Penerima, error) {
	var penerima models.Penerima
	var db = database.DB

	// penerima.IDSurat = arg.SuratMasuk.ID

	// arg.Penerima.IDSurat = arg.SuratMasuk.ID

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreatePenerimaSuratMasuk,
		arg.Penerima.IDSurat,
		arg.Penerima.IDPengguna,
		arg.Penerima.CreatedAt2).StructScan(&penerima)
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

type UpdateSuratMasukParams struct {
	SuratMasuk models.SuratMasuk
	Penerima   models.Penerima
}

func (r *repo) Update(arg UpdateSuratMasukParams) (*models.CreateSuratMasuk, error) {

	var suratMasuk models.CreateSuratMasuk
	var err error

	var surat *models.SuratMasuk
	surat, err = r.updateSurat(&arg)
	if err != nil {
		return nil, err
	}

	suratMasuk.SuratMasuk = *surat

	var penerima *models.Penerima
	arg.Penerima.IDSurat = surat.ID
	arg.Penerima.CreatedAt2 = surat.CreatedAt
	penerima, err = r.updatePenerima(&arg)
	if err != nil {
		return nil, err
	}

	suratMasuk.Penerima = *penerima

	return &suratMasuk, nil
}

func (*repo) updateSurat(arg *UpdateSuratMasukParams) (*models.SuratMasuk, error) {
	var suratMasuk models.SuratMasuk
	var db = database.DB
	var err error

	var byteUpload []byte
	var filename string
	var fullpath string

	uploadPayload := arg.SuratMasuk.Upload
	if uploadPayload != nil {
		str := strings.SplitAfter(*uploadPayload, ",")
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

		filename = "surat_masuk_" + time.Now().Format(helper.LayoutTime3) + "." + extFile
		fullpath = "http://" + os.Getenv("ftp_addr") + ":" + os.Getenv("ftp_port_image") + "/" + filename
	}

	arg.SuratMasuk.Upload = &fullpath

	tx := db.MustBegin()
	err = tx.QueryRowx(query.UpdateSuratMasuk,
		arg.SuratMasuk.Tanggal,
		arg.SuratMasuk.Nomor,
		arg.SuratMasuk.IDPengirim,
		arg.SuratMasuk.Perihal,
		arg.SuratMasuk.IDJenis,
		arg.SuratMasuk.Keterangan,
		arg.SuratMasuk.UpdatedAt,
		arg.SuratMasuk.Upload,
		arg.SuratMasuk.ID,
	).StructScan(&suratMasuk)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	err = helper.Upload(byteUpload, filename)
	if err != nil {
		return nil, err
	}

	return &suratMasuk, nil
}

func (*repo) updatePenerima(arg *UpdateSuratMasukParams) (*models.Penerima, error) {
	var penerima models.Penerima
	var db = database.DB
	var err error

	_, err = db.Exec("DELETE FROM tbl_penerima WHERE id_surat = $1", arg.SuratMasuk.ID)
	if err != nil {
		return nil, err
	}

	tx := db.MustBegin()
	err = tx.QueryRowx(query.UpdatePenerimaSuratMasuk,
		arg.Penerima.IDSurat,
		arg.Penerima.IDPengguna,
		arg.Penerima.CreatedAt2,
		arg.Penerima.UpdatedAt2).StructScan(&penerima)
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
