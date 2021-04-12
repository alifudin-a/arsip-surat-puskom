package models

// SuratKeluar : struct untuk create dan update
type SuratKeluar struct {
	ID         int64  `json:"id" db:"id"`
	Tanggal    string `json:"tanggal" db:"tanggal"`
	Nomor      string `json:"nomor" db:"nomor"`
	IDPengirim int64  `json:"id_pengirim" db:"id_pengirim"`
	Perihal    string `json:"perihal" db:"perihal"`
	IDJenis    int64  `json:"id_jenis" db:"id_jenis"`
	Keterangan string `json:"keterangan" db:"keterangan"`
}

// PenerimaSuratKeluar : struct untuk create dan update
// type PenerimaSuratKeluar struct {
// }

type SelectSuratKeluar struct {
	ID         int64   `json:"id" db:"id"`
	Tanggal    string  `json:"tanggal" db:"tanggal" validate:"required"`
	Nomor      string  `json:"nomor" db:"nomor" validate:"required"`
	IDPengirim int64   `json:"id_pengirim,omitempty" db:"id_pengirim"`
	Pengirim   string  `json:"pengirim,omitempty" db:"pengirim"`
	Perihal    string  `json:"perihal" db:"perihal" validate:"required"`
	IDJenis    *int64  `json:"id_jenis,omitempty" db:"id_jenis"`
	Jenis      *string `json:"jenis,omitempty" db:"jenis"`
	Keterangan *string `json:"keterangan,omitempty" db:"keterangan"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}

type ListPenerimaSuratKeluar struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type ListSuratKeluar struct {
	SelectSuratKeluar
	PenerimaSuratKeluar []ListPenerimaSuratKeluar `json:"penerima"`
}

type ReadSuratKeluar struct {
	SelectSuratKeluar
	PenerimaSuratKeluar []ListPenerimaSuratKeluar `json:"penerima"`
}
