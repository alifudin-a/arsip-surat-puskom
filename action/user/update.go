package action

import (
	"net/http"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/user"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/user"
	"github.com/labstack/echo/v4"
)

// func validateUpdate(req *models.User, c echo.Context) (err error) {

// 	if err = c.Bind(req); err != nil {
// 		return err
// 	}

// 	return c.Validate(req)
// }

func UpdateUserHandler(c echo.Context) (err error) {

	var resp helper.Response
	var req = new(models.User)
	var user *models.User

	if err = c.Bind(&req); err != nil {
		return
	}

	repo := repository.NewUserRepository()

	arg := repository.UpdateUserParams{
		ID:        req.ID,
		Name:      req.Name,
		Fullname:  req.FullName,
		UpdatedAt: *req.UpdatedAt,
	}

	exist, err := repo.IsExist(repository.IsExistUserParams{
		ID: int64(arg.ID),
	})
	if err != nil {
		return err
	}

	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)

	}

	user, err = repo.Update(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil mebgubah data!"
	resp.Body = map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusOK, resp)
}
