package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
)

func CreateSuratMasuk(params *models.CreateSuratMasuk) repository.CreateSuratMasukParams {
	var res repository.CreateSuratMasukParams

	t := time.Now()
	res.SuratMasuk.Tanggal = t.Format(helper.LayoutTime2)
	res.SuratMasuk.Nomor = params.Nomor
	res.SuratMasuk.IDPengirim = params.IDPengirim
	res.SuratMasuk.Perihal = params.Perihal
	res.SuratMasuk.IDJenis = params.IDJenis
	res.SuratMasuk.Keterangan = params.Keterangan
	res.SuratMasuk.CreatedAt = helper.NullString(t.Format(helper.LayoutTime))

	res.Penerima.IDSurat = params.IDSurat
	res.Penerima.IDPengguna = params.IDPengguna
	res.Penerima.CreatedAt2 = helper.NullString(t.Format(helper.LayoutTime))

	return res
}

func UpdateSuratMasuk(params *models.CreateSuratMasuk) repository.UpdateSuratMasukParams {
	var res repository.UpdateSuratMasukParams

	t := time.Now()
	res.SuratMasuk.ID = params.ID
	res.SuratMasuk.Tanggal = t.Format(helper.LayoutTime2)
	res.SuratMasuk.Nomor = params.Nomor
	res.SuratMasuk.IDPengirim = params.IDPengirim
	res.SuratMasuk.Perihal = params.Perihal
	res.SuratMasuk.IDJenis = params.IDJenis
	res.SuratMasuk.Keterangan = params.Keterangan
	res.SuratMasuk.UpdatedAt = helper.NullString(t.Format(helper.LayoutTime))

	res.Penerima.IDSurat = params.IDSurat
	res.Penerima.IDPengguna = params.IDPengguna
	res.Penerima.UpdatedAt2 = helper.NullString(t.Format(helper.LayoutTime))

	return res
}
