package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/user"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/user"
	"github.com/labstack/echo/v4"
)

func ListUserHandler(c echo.Context) (err error) {

	var user []models.User
	var resp helper.Response

	repo := repository.NewUserRepository()

	user, err = repo.FindAll()
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menampilkan data!"
	resp.Body = map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusOK, resp)
}
