package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
)

func CreateSuratKeluar(params *models.CreateSuratKeluar) repository.CreateSuratKeluarParams {
	var res repository.CreateSuratKeluarParams

	t := time.Now()
	res.SuratKeluar.Nomor = params.Nomor
	res.SuratKeluar.IDPengirim = params.IDPengirim
	res.SuratKeluar.Perihal = params.Perihal
	res.SuratKeluar.IDJenis = params.IDJenis
	res.SuratKeluar.Keterangan = params.Keterangan
	res.SuratKeluar.CreatedAt = helper.NullString(t.Format(helper.LayoutTime))

	res.PenerimaSuratKeluar = params.PenerimaSuratKeluar

	return res
}

func UpdateSuratKeluar(params *models.CreateSuratKeluar) repository.UpdateSuratKeluarParams {
	var res repository.UpdateSuratKeluarParams

	t := time.Now()

	res.SuratKeluar.ID = params.ID
	res.SuratKeluar.Tanggal = params.Tanggal
	res.SuratKeluar.Nomor = params.Nomor
	res.SuratKeluar.IDPengirim = params.IDPengirim
	res.SuratKeluar.Perihal = params.Perihal
	res.SuratKeluar.IDJenis = params.IDJenis
	res.SuratKeluar.Keterangan = params.Keterangan
	res.SuratKeluar.UpdatedAt = helper.NullString(t.Format(helper.LayoutTime))

	res.PenerimaSuratKeluar = params.PenerimaSuratKeluar

	return res
}
