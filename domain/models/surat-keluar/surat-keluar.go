package models

// SuratKeluar model struct
type SuratKeluar struct {
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

type ListSuratKeluar struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	Pengirim   string  `json:"pengirim" db:"pengirim"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	Jenis      *string `json:"jenis" db:"jenis"`
	Keterangan *string `json:"keterangan" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}

// SuratKeluar model struct
type CreateSuratKeluar struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	IDPengirim int64   `json:"id_pengirim" db:"id_pengirim" validate:"required"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64  `json:"-" db:"id_jenis"`
	Keterangan *string `json:"keterangan" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}

type SelectJoinSuratKeluar struct {
	ID         int64              `json:"id" db:"id"`
	Tanggal    string             `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string             `json:"nomor" db:"nomor" validate:"required"`
	Pengirim   string             `json:"pengirim" db:"pengirim"`
	Perihal    string             `json:"perihal" db:"perihal" validate:"required"`
	Keterangan *string            `json:"keterangan" db:"keterangan"`
	Penerima   SelectJoinPenerima `json:"penerima"`
}

type SelectJoinPenerima struct {
	ID       int64  `json:"id"`
	IDSurat  int64  `json:"id_surat"`
	Penerima string `json:"name"`
}
