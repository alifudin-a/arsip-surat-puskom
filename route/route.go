package route

import (
	"net/http"

	actionJenisSurat "github.com/alifudin-a/arsip-surat-puskom/action/jenis-surat"
	actionLogin "github.com/alifudin-a/arsip-surat-puskom/action/login"
	actionPenerima "github.com/alifudin-a/arsip-surat-puskom/action/penerima"
	action "github.com/alifudin-a/arsip-surat-puskom/action/surat"
	actionSurat "github.com/alifudin-a/arsip-surat-puskom/action/surat"
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

	login := actionLogin.NewLoginHandler()
	api.POST("/login", login.LoginHandler)

	pList := actionPenerima.NewListPenerima()
	pRead := actionPenerima.NewReadPenerima()
	pDelete := actionPenerima.NewDeletePenerima()
	pCreate := actionPenerima.NewCreatePenerima()
	pUpdate := actionPenerima.NewUpdatePenerima()

	api.GET("/penerima", pList.ListPenerimaHandler)
	api.GET("/penerima/:id", pRead.ReadPenerimaHandler)
	api.DELETE("/penerima/:id", pDelete.DeletePenerimaHandler)
	api.POST("/penerima", pCreate.CreatePenerimaHandler)
	api.PUT("/penerima", pUpdate.UpdatePenerimaHandler)

	sList := actionSurat.NewListSurat()
	sRead := actionSurat.NewReadSurat()
	sDelete := actionSurat.NewDeleteSurat()
	sCreate := actionSurat.NewCreateSurat()
	sUpdate := actionSurat.NewUpdateSurat()

	api.GET("/surat", sList.ListSuratHandler)
	api.GET("/surat/:id", sRead.ReadSuratHandler)
	api.DELETE("/surat/:id", sDelete.DeleteSuratHandler)
	api.POST("/surat", sCreate.CreateSuratHandler)
	api.PUT("/surat", sUpdate.UpdateSuratHandler)

	api.POST("/surat2", actionSurat.NewCreate2Surat().CreateSurat2Handler)
	api.POST("/surat3", action.NewCreate3Surat().CreateSurat3Handler)

	e.Logger.Fatal(e.Start(":9000"))

	return e
}
