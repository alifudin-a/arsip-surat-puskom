package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/login"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type LoginRepository interface {
	Login(arg LoginParams) (*models.Login, error)
}

type repo struct{}

func NewLoginRepository() LoginRepository {
	return &repo{}
}

type LoginParams struct {
	Username string
	Password string
}

func (*repo) Login(arg LoginParams) (*models.Login, error) {
	var login models.Login
	var db = database.DB

	err := db.Get(&login, query.Login, arg.Username, arg.Password)
	if err != nil {
		return nil, err
	}

	return &login, nil
}
