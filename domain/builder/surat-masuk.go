package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/surat-masuk"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/surat-masuk"
)

func CreaetSuratMasuk(params *models.CreateSuratMasuk) repository.CreateSuratMasukParams {
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
