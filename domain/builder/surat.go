package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
)

func CreateSurat(params *models.XCreateSurat) repository.CreateSurat {
	var res repository.CreateSurat

	t := time.Now()

	res.Surat.Tanggal = params.Tanggal
	res.Surat.Nomor = params.Nomor
	res.Surat.IDPengirim = params.IDPengirim
	res.Surat.Perihal = params.Perihal
	res.Surat.IDJenis = params.IDJenis
	res.Surat.Keterangan = params.Keterangan
	res.Surat.CreatedAt = t.Format(helper.LayoutTime)

	res.Penerima = params.Penerima

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
