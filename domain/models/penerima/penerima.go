package models

type Penerima struct {
	ID         int64   `json:"id" db:"id"`
	IDSurat    int64   `json:"id_surat" db:"id_surat"`
	IDPengguna int64   `json:"id_pengguna" db:"id_pengguna"`
	Name       string  `json:"name,omitempty" db:"name"`
	CreatedAt  *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}
