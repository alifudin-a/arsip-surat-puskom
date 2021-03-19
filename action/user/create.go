package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/user"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/user"
	"github.com/labstack/echo/v4"
)

type Create struct{}

func NewCreateUser() *Create {
	return &Create{}
}

func (cr *Create) validate(req *models.User, c echo.Context) (err error) {

	if err = c.Bind(req); err != nil {
		return err
	}

	return c.Validate(req)
}

func (cr *Create) CreateUserHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.User)
	var user *models.User

	err = cr.validate(req, c)
	if err != nil {
		return err
	}

	repo := repository.NewUserRepository()

	arg := repository.CreateUserParams{
		Name:      req.Name,
		Fullname:  req.FullName,
		CreatedAt: *req.CreatedAt,
	}

	user, err = repo.Create(arg)
	if err != nil {
		return
	}

	resp.Code = http.StatusCreated
	resp.Message = "Berhasil membuat user!"
	resp.Body = map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusCreated, resp)
}
