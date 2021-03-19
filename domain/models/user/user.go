package models

// User model struct
type User struct {
	ID        int64   `json:"id" db:"id"`
	Name      string  `json:"name" db:"name" validate:"required"`
	FullName  string  `json:"fullname" db:"fullname" validate:"required"`
	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
}
