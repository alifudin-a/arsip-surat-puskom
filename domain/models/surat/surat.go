package models

// SuratMasuk model struct
type Surat struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	IDPenerima int64   `json:"id_penerima" db:"id_penerima" validate:"required"`
	IDPengirim int64   `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64  `json:"id_jenis" db:"id_jenis"`
	Keterangan *string `json:"keterangan" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}

type ListSurat struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	Penerima   string  `json:"penerima" db:"penerima"`
	Pengirim   string  `json:"pengirim" db:"pengirim"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	Jenis      *string `json:"jenis" db:"jenis"`
	Keterangan *string `json:"keterangan" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}
