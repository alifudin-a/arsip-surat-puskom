package models

// SuratMasuk model struct
type Surat struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	IDPengirim int64   `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64  `json:"id_jenis" db:"id_jenis"`
	Keterangan *string `json:"keterangan" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}

type SelectSurat struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	IDPengirim int64   `json:"id_pengirim,omitempty" db:"id_pengirim"`
	Pengirim   string  `json:"pengirim,omitempty" db:"pengirim"`
	Penerima   string  `json:"penerima,omitempty" db:"penerima"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64  `json:"id_jenis,omitempty" db:"id_jenis"`
	Jenis      *string `json:"jenis,omitempty" db:"jenis"`
	Keterangan *string `json:"keterangan,omitempty" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}

type Penerima struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"penerima" db:"penerima"`
}

type ListSurat struct {
	SelectSurat
	PenerimaSurat []Penerima `json:"penerima"`
}
