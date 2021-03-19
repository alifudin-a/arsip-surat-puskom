package repository

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/user"
	"github.com/alifudin-a/arsip-surat-puskom/domain/query"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(arg ReadUserParams) (*models.User, error)
	Delete(arg DeleteUserParams) (err error)
	IsExist(arg IsExistUserParams) (bool, error)
	Create(arg CreateUserParams) (*models.User, error)
	Update(arg UpdateUserParams) (*models.User, error)
}

type repo struct{}

func NewUserRepository() UserRepository {
	return &repo{}
}

// ListUserParams
// type ListUser struct{}

func (*repo) FindAll() ([]models.User, error) {

	var user []models.User

	var db = database.OpenDB()

	err := db.Select(&user, query.ListAllUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ReadUserlParams .
type ReadUserParams struct {
	ID int64
}

func (*repo) FindById(arg ReadUserParams) (*models.User, error) {
	var user models.User
	var db = database.OpenDB()

	err := db.Get(&user, query.ReadUserByID, arg.ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUserParams .
type DeleteUserParams struct {
	ID int64
}

func (*repo) Delete(arg DeleteUserParams) (err error) {
	var db = database.OpenDB()

	tx := db.MustBegin()
	_, err = tx.Exec(query.DeleteUser, arg.ID)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return nil
}

// IsExistUserParams .
type IsExistUserParams struct {
	ID int64
}

func (*repo) IsExist(arg IsExistUserParams) (bool, error) {
	var db = database.OpenDB()
	var total int

	err := db.Get(&total, query.IsExistUser, arg.ID)
	if err != nil {
		return false, nil
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

// CreateUserParams .
type CreateUserParams struct {
	Name      string
	Fullname  string
	CreatedAt string
}

func (*repo) Create(arg CreateUserParams) (*models.User, error) {
	var user models.User
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.CreateUser, arg.Name, arg.Fullname, arg.CreatedAt).StructScan(&user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserParams .
type UpdateUserParams struct {
	Name      string
	Fullname  string
	UpdatedAt string
	ID        int64
}

func (*repo) Update(arg UpdateUserParams) (*models.User, error) {
	var user models.User
	var db = database.OpenDB()

	tx := db.MustBegin()
	err := tx.QueryRowx(query.UpdateUser, arg.Name, arg.Fullname, arg.UpdatedAt, arg.ID).StructScan(&user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
