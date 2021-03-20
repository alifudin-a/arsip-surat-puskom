package models

// SuratKeluar model struct
type SuratKeluar struct {
	ID         int64  `json:"id" db:"id" validate:"required"`
	Tanggal    string `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      int64  `json:"nomor" db:"nomor" validate:"required"`
	IDPenerima int64  `json:"id_penerima" db:"id_penerima" validate:"required"`
	IDPengirim int64  `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Perihal    string `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    string `json:"id_jenis" db:"id_jenis"`
	Keterangan string `json:"keterangan" db:"keterangan"`
}

type ReadSuratKeluar struct {
	ID         int64  `json:"id" db:"id" validate:"required"`
	Tanggal    string `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      int64  `json:"nomor" db:"nomor" validate:"required"`
	IDPenerima int64  `json:"id_penerima" db:"id_penerima" validate:"required"`
	Penerima   string `json:"penerima" db:"penerima"`
	IDPengirim int64  `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Pengirim   string `json:"pengirim" db:"pengirim"`
	Perihal    string `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    string `json:"id_jenis" db:"id_jenis"`
	Jenis      string `json:"jenis" db:"jenis"`
	Keterangan string `json:"keterangan" db:"keterangan"`
}
