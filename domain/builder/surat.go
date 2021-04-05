package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat"
)

func CreateSurat(params *models.CreateSuratPenerima) repository.CreateSurat {
	var res repository.CreateSurat

	t := time.Now()

	res.Surat.Tanggal = params.Tanggal
	res.Surat.Nomor = params.Nomor
	res.Surat.IDPengirim = params.IDPengirim
	res.Surat.Perihal = params.Perihal
	res.Surat.IDJenis = params.IDJenis
	res.Surat.Keterangan = params.Keterangan
	res.Surat.CreatedAt = helper.NullString(t.Format(helper.LayoutTime))

	res.Penerima = params.Penerima

	return res
}

func UpdateSurat(params *models.CreateSuratPenerima) repository.UpdateSuratParams {
	var res repository.UpdateSuratParams

	t := time.Now()

	res.Surat.ID = params.ID
	res.Surat.Tanggal = params.Tanggal
	res.Surat.Nomor = params.Nomor
	res.Surat.IDPengirim = params.IDPengirim
	res.Surat.Perihal = params.Perihal
	res.Surat.IDJenis = params.IDJenis
	res.Surat.Keterangan = params.Keterangan
	res.Surat.UpdatedAt = helper.NullString(t.Format(helper.LayoutTime))

	res.Penerima = params.Penerima

	return res
}
