package route

import (
	"net/http"

	actionJenisSurat "github.com/alifudin-a/arsip-surat-puskom/action/jenis-surat"
	actionLogin "github.com/alifudin-a/arsip-surat-puskom/action/login"
	actionPenerima "github.com/alifudin-a/arsip-surat-puskom/action/penerima"
	actionSurat "github.com/alifudin-a/arsip-surat-puskom/action/surat"
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

	// validator json body
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = helper.CustomReadableError

	// middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "==> METHOD=${method}, URI=${uri}, STATUS=${status}, " +
			"HOST=${host}, ERROR=${error}, LATENCY_HUMAN=${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "api-key", "user-token"},
		ExposeHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentLength, echo.HeaderAcceptEncoding, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderContentDisposition, "api-key", "user-token"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	// middleware validasi token jwt
	// e.Use(validation.Validation())

	// e.POST("/login", actionLogin.NewLoginHandler().LoginHandler)

	// endpoint group
	api := e.Group("/api")

	// api.Use(validation.Validation())

	// endpoint user / pengguna
	api.GET("/user", actionUser.NewListUser().ListUserHandler)
	api.GET("/user/:id", actionUser.NewReadUser().ReadUserHandler)
	api.DELETE("/user/:id", actionUser.NewDeleteUser().DeleteUserHandler)
	api.POST("/user", actionUser.NewCreateUser().CreateUserHandler)
	api.PUT("/user", actionUser.NewUpdateUser().UpdateUserHandler)

	// endpoint jenis_surat
	api.GET("/jenis_surat", actionJenisSurat.NewListJenisSurat().ListJenisSuratHandler)
	api.GET("/jenis_surat/:id", actionJenisSurat.NewReadJenisSurat().ReadJenisSuratHandler)
	api.DELETE("/jenis_surat/:id", actionJenisSurat.NewDeleteJenisSurat().DeleteJenisSuratHandler)
	api.POST("/jenis_surat", actionJenisSurat.NewCreateJenisSurat().CreateJenisSuratHandler)
	api.PUT("/jenis_surat", actionJenisSurat.NewUpdateJenisSurat().UpdateUserHandler)

	// endpoint login
	api.POST("/login", actionLogin.NewLoginHandler().LoginHandler)

	// endpoint penerima
	api.GET("/penerima", actionPenerima.NewListPenerima().ListPenerimaHandler)
	api.GET("/penerima/:id", actionPenerima.NewReadPenerima().ReadPenerimaHandler)
	api.DELETE("/penerima/:id", actionPenerima.NewDeletePenerima().DeletePenerimaHandler)
	api.POST("/penerima", actionPenerima.NewCreatePenerima().CreatePenerimaHandler)
	api.PUT("/penerima", actionPenerima.NewUpdatePenerima().UpdatePenerimaHandler)

	// endpoint surat
	api.GET("/surat", actionSurat.NewListSurat().ListSuratHandler)
	api.GET("/surat/:id", actionSurat.NewReadSurat().ReadSuratHandler)
	api.DELETE("/surat/:id", actionSurat.NewDeleteSurat().DeleteSuratHandler)
	api.PUT("/surat", actionSurat.NewUpdateSurat().UpdateSuratHandler)
	api.POST("/surat", actionSurat.NewCreateSurat().CreateSuratHandler)

	// endpoint surat masuk
	api.GET("/penerima_surat_masuk/:id", actionSuratMasuk.NewListSuratMasuk().ListPenerimaSuratMasukHandler)
	api.GET("/surat_masuk", actionSuratMasuk.NewListSuratMasuk().ListSuratMasukHandler)
	api.GET("/surat_masuk/:id", actionSuratMasuk.NewReadSuratMasuk().ReadSuratMasukHandler)
	api.DELETE("/surat_masuk/:id", actionSuratMasuk.NewDeleteSuratMasuk().DeleteSuratMasukHandler)
	api.POST("/surat_masuk", actionSuratMasuk.NewCreateSuratMasuk().CreateSuratMasukHandler)
	api.PUT("/surat_masuk", actionSuratMasuk.NewUpdateSuratMasuk().UpdateSuratMasukHandler)

	// endpoint surat keluar
	api.GET("/pengirim_surat_keluar/:id", actionSuratKeluar.NewListSuratKeluar().ListSuratKeluarByIDPengirim)
	api.GET("/surat_keluar", actionSuratKeluar.NewListSuratKeluar().ListSuratKeluarHandler)
	api.GET("/surat_keluar/:id", actionSuratKeluar.NewReadSuratKeluar().ReadSuratKeluarHandler)
	api.DELETE("/surat_keluar/:id", actionSuratKeluar.NewDeleteSuratKeluar().DeleteSuratKeluarHandler)
	api.POST("/surat_keluar", actionSuratKeluar.NewCreateSuratKeluar().CreateSuratKeluarHandler)
	api.PUT("/surat_keluar", actionSuratKeluar.NewUpdateSuratKeluar().UpdateSuratHandler)

	e.Logger.Fatal(e.Start(":9000"))

	return e
}
