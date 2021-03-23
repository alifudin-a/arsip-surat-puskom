package models

type Login struct {
	ID         int     `json:"id" db:"id"`
	IDPengguna int     `json:"id_pengguna" db:"id_pengguna"`
	Username   string  `json:"username" db:"username"`
	Password   string  `json:"password,omitempty" db:"password"`
	CreatedAt  *string `json:"created_at" db:"created_at"`
	UpdatedAt  *string `json:"updated_at,omitempty" db:"updated_at"`
}
