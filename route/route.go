package route

import (
	"net/http"

	actionJenisSurat "github.com/alifudin-a/arsip-surat-puskom/action/jenis-surat"
	actionUser "github.com/alifudin-a/arsip-surat-puskom/action/user"
	"github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helper.CustomReadableError

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, URI=${uri}, STATUS=${status}, " +
			"HOST=${host}, ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("/api")

	usrCreate := actionUser.NewCreateUser()
	usrDelete := actionUser.NewDeleteUser()
	usrList := actionUser.NewListUser()
	usrRead := actionUser.NewReadUser()
	usrUpdate := actionUser.NewUpdateUser()

	api.GET("/user", usrList.ListUserHandler)
	api.GET("/user/:id", usrRead.ReadUserHandler)
	api.DELETE("/user/:id", usrDelete.DeleteUserHandler)
	api.POST("/user", usrCreate.CreateUserHandler)
	api.PUT("/user", usrUpdate.UpdateUserHandler)

	jsCreate := actionJenisSurat.NewCreateJenisSurat()
	jsDelete := actionJenisSurat.NewDeleteJenisSurat()
	jsList := actionJenisSurat.NewListJenisSurat()
	jsRead := actionJenisSurat.NewReadJenisSurat()
	jsUpdate := actionJenisSurat.NewUpdateJenisSurat()

	api.GET("/jenis_surat", jsList.ListJenisSuratHandler)
	api.GET("/jenis_surat/:id", jsRead.ReadJenisSuratHandler)
	api.DELETE("/jenis_surat/:id", jsDelete.DeleteJenisSuratHandler)
	api.POST("/jenis_surat", jsCreate.CreateJenisSuratHandler)
	api.PUT("/jenis_surat", jsUpdate.UpdateUserHandler)

	e.Logger.Fatal(e.Start(":9000"))

	return e
}
