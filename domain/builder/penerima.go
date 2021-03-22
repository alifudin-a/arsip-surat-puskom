package builder

import (
	"time"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/penerima"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/penerima"
)

func CreatePenerima(params *models.Penerima) repository.CreatePenerimaParams {
	var res repository.CreatePenerimaParams

	t := time.Now()

	res.IDSurat = params.IDSurat
	res.IDPengguna = params.IDPengguna
	res.CreatedAt = t.Format(helper.LayoutTime)

	return res
}

func UpdatePenerima(params *models.Penerima) repository.UpdatePenerimaParams {
	var res repository.UpdatePenerimaParams

	t := time.Now()

	res.ID = params.ID
	res.IDSurat = params.IDSurat
	res.IDPengguna = params.IDPengguna
	res.UpdatedAt = t.Format(helper.LayoutTime)

	return res
}
