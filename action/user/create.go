package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/user"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/user"
	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.User)
	var user *models.User

	if err = c.Bind(&req); err != nil {
		return
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
