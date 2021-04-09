package models

import "github.com/alifudin-a/arsip-surat-puskom/domain/helper"

type SuratMasuk struct {
	ID         int64             `json:"id" db:"id"`
	Tanggal    string            `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string            `json:"nomor" db:"nomor" validate:"required"`
	IDPengirim int64             `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Perihal    string            `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64            `json:"id_jenis" db:"id_jenis"`
	Keterangan *string           `json:"keterangan" db:"keterangan"`
	CreatedAt  helper.NullString `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  helper.NullString `json:"updated_at,omitempty" db:"updated_at"`
}

type ListSuratMasuk struct {
	ID         int64             `json:"id" db:"id"`
	Tanggal    string            `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string            `json:"nomor" db:"nomor" validate:"required"`
	IDPengirim int64             `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Pengirim   string            `json:"pengirim" db:"pengirim"`
	Perihal    string            `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64            `json:"id_jenis" db:"id_jenis"`
	Jenis      *string           `json:"jenis" db:"jenis"`
	Keterangan *string           `json:"keterangan" db:"keterangan"`
	CreatedAt  helper.NullString `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  helper.NullString `json:"updated_at,omitempty" db:"updated_at"`
}
