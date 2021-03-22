package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-keluar"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-keluar"
)

func CreateSuratKeluar(params *models.SuratKeluar) repository.CreateSuratKeluarParams {
	var res repository.CreateSuratKeluarParams

	t := time.Now()

	res.Tanggal = params.Tanggal
	res.Nomor = params.Nomor
	res.IDPengirim = params.IDPengirim
	res.Perihal = params.Perihal
	res.IDJenis = *params.IDJenis
	res.Keterangan = *params.Keterangan
	res.CreatedAt = t.Format(helper.LayoutTime)

	return res
}

func UpdateSuratKeluar(params *models.SuratKeluar) repository.UpdateSuratKeluarParams {
	var res repository.UpdateSuratKeluarParams

	t := time.Now()

	res.ID = params.ID
	res.Tanggal = params.Tanggal
	res.Nomor = params.Nomor
	res.IDPengirim = params.IDPengirim
	res.Perihal = params.Perihal
	res.IDJenis = *params.IDJenis
	res.Keterangan = *params.Keterangan
	res.UpdatedAt = t.Format(helper.LayoutTime)

	return res
}

func CreateSuratKeluarV2(params *models.CreateSuratKeluar) repository.CreateSuratKeluarParamsV2 {
	var res repository.CreateSuratKeluarParamsV2

	t := time.Now()

	res.Tanggal = params.Tanggal
	res.Nomor = params.Nomor
	res.IDPengirim = params.IDPengirim
	res.Perihal = params.Perihal
	res.Keterangan = *params.Keterangan
	res.CreatedAt = t.Format(helper.LayoutTime)

	return res
}
