package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	models "github.com/alifudin-a/arsip-surat-puskom/domain/models/user"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/user"
	"github.com/labstack/echo/v4"
)

type Read struct{}

func NewReadUser() *Read {
	return &Read{}
}

func (rd *Read) ReadUserHandler(c echo.Context) (err error) {
	var resp helper.Response

	var user = &models.User{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Message = "ID harus berupa angka!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	repo := repository.NewUserRepository()

	arg := repository.ReadUserParams{
		ID: int64(id),
	}

	user, err = repo.FindById(arg)
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
