package route

import (
	"net/http"

	actionJenisSurat "github.com/alifudin-a/arsip-surat-puskom/action/jenis-surat"
	actionLogin "github.com/alifudin-a/arsip-surat-puskom/action/login"
	actionSuratKeluar "github.com/alifudin-a/arsip-surat-puskom/action/surat-keluar"
	actionSuratMasuk "github.com/alifudin-a/arsip-surat-puskom/action/surat-masuk"
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

	skList := actionSuratKeluar.NewListSuratKeluar()
	skRead := actionSuratKeluar.NewReadSuratKeluar()
	skDelete := actionSuratKeluar.NewDeleteSuratKeluar()
	skCreate := actionSuratKeluar.NewCreateSuratKeluar()
	skUpdate := actionSuratKeluar.NewUpdateSuratKeluar()

	api.GET("/surat_keluar", skList.ListSuratKeluarHandler)
	api.GET("/surat_keluar/:id", skRead.ReadSuratKeluarHandler)
	api.DELETE("/surat_keluar/:id", skDelete.DeleteSuratKeluarHandler)
	api.POST("/surat_keluar", skCreate.NewCreateSuratKeluarHandler)
	api.PUT("/surat_keluar", skUpdate.NewUpdateSuratKeluarHandler)

	smList := actionSuratMasuk.NewListSuratMasuk()
	smRead := actionSuratMasuk.NewReadSuratMasuk()
	smDelete := actionSuratMasuk.NewDeleteSuratMasuk()
	smCreate := actionSuratMasuk.NewCreateSuratMasuk()
	smUpdate := actionSuratMasuk.NewUpdateSuratMasuk()

	api.GET("/surat_masuk", smList.ListSuratMasukHandler)
	api.GET("/surat_masuk/:id", smRead.ReadSuratMasukHandler)
	api.DELETE("/surat_masuk/:id", smDelete.DeleteSuratMasukHandler)
	api.POST("/surat_masuk", smCreate.NewCreateSuratMasukHandler)
	api.PUT("/surat_masuk", smUpdate.NewUpdateSuratMasukHandler)

	login := actionLogin.NewLoginHandler()
	loginV2 := actionLogin.NewLoginHandlerV2()

	api.POST("/login", login.LoginHandler)
	api.POST("/login_v2", loginV2.LoginHandlerV2)

	skCreateV2 := actionSuratKeluar.NewCreateSuratKeluarV2()
	api.POST("/surat_keluar_v2", skCreateV2.NewCreateSuratKeluarHandlerV2)

	e.Logger.Fatal(e.Start(":9000"))

	return e
}
