package models

// JenisSurat model struct
type JenisSurat struct {
	ID        int64   `json:"id" db:"id"`
	Kode      int64   `json:"kode" db:"kode"`
	Name      string  `json:"name" db:"name"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
}
