package models

// JabatanStruktural model struct
type JabatanStruktural struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name" validate:"required"`
	FullName string `json:"fullname" db:"fullname" validate:"required"`
}
