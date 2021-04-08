package action

import (
	"net/http"
	"strconv"

	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	repository "github.com/alifudin-a/arsip-surat-puskom/repository/user"
	mid "github.com/alifudin-a/arsip-surat-puskom/route/middleware"
	"github.com/labstack/echo/v4"
)

type Delete struct{}

func NewDeleteUser() *Delete {
	return &Delete{}
}

func (dl *Delete) DeleteUserHandler(c echo.Context) (err error) {
	var resp helper.Response

	err = mid.ValidationKey(c)
	if err != nil {
		return
	}

	err = mid.ValidationJWT(c)
	if err != nil {
		return
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	repo := repository.NewUserRepository()

	arg := repository.DeleteUserParams{
		ID: int64(id),
	}

	exist, err := repo.IsExist(repository.IsExistUserParams{ID: id})
	if err != nil {
		return err
	}

	if !exist {
		resp.Code = http.StatusBadRequest
		resp.Message = "Data tidak ada!"
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = repo.Delete(arg)
	if err != nil {
		return err
	}

	resp.Code = http.StatusOK
	resp.Message = "Berhasil menghapus data!"

	return c.JSON(http.StatusOK, resp)
}
