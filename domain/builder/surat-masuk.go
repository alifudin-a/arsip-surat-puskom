package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
)

func CreateSuratMasuk(params *models.SuratMasuk) repository.CreateSuratMasukParams {
	var res repository.CreateSuratMasukParams

	t := time.Now()

	res.Tanggal = params.Tanggal
	res.Nomor = params.Nomor
	res.IDPenerima = params.IDPenerima
	res.IDPengirim = params.IDPengirim
	res.Perihal = params.Perihal
	res.IDJenis = *params.IDJenis
	res.Keterangan = *params.Keterangan
	res.CreatedAt = t.Format(helper.LayoutTime)

	return res
}

func UpdateSuratMasuk(params *models.SuratMasuk) repository.UpdateSuratMasukParams {
	var res repository.UpdateSuratMasukParams

	t := time.Now()

	res.ID = params.ID
	res.Tanggal = params.Tanggal
	res.Nomor = params.Nomor
	res.IDPenerima = params.IDPenerima
	res.IDPengirim = params.IDPengirim
	res.Perihal = params.Perihal
	res.IDJenis = *params.IDJenis
	res.Keterangan = *params.Keterangan
	res.UpdatedAt = t.Format(helper.LayoutTime)

	return res
}
