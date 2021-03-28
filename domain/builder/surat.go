package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
)

func CreateSurat(params *models.Surat) repository.CreateSuratParams {
	var res repository.CreateSuratParams

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

func CreateSurat2(params *models.CreateSurat) repository.CreateSurat2Params {
	var res repository.CreateSurat2Params

	t := time.Now()

	res.Tanggal = params.Tanggal
	res.Nomor = params.Nomor
	res.IDPengirim = params.IDPengirim
	res.Perihal = params.Perihal
	res.IDJenis = *params.IDJenis
	res.Keterangan = *params.Keterangan
	res.CreatedAt = t.Format(helper.LayoutTime)
	res.IDSurat = params.IDSurat
	res.IDPengguna = params.IDPengguna
	res.CreatedAt2 = t.Format(helper.LayoutTime)

	return res
}

func UpdateSurat(params *models.Surat) repository.UpdateSuratParams {
	var res repository.UpdateSuratParams

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
