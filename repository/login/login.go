package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/login"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type LoginRepository interface {
	GetUsername(arg GetUsernameParams) (*models.Login, error)
}

type repo struct{}

func NewLoginRepository() LoginRepository {
	return &repo{}
}

type GetUsernameParams struct {
	Username string
	Password string
}

func (*repo) GetUsername(arg GetUsernameParams) (*models.Login, error) {
	var login models.Login
	var db = database.OpenDB()

	err := db.Get(&login, query.GetUsername, arg.Username, arg.Password)
	if err != nil {
		return nil, err
	}

	return &login, nil
}
