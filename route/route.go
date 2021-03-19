package route

import (
	"net/http"

	actionJenisSurat "github.com/alifudin-a/arsip-surat-puskom/action/jenis-surat"
	actionUser "github.com/alifudin-a/arsip-surat-puskom/action/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute() *echo.Echo {
	e := echo.New()

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
	api.GET("/user", actionUser.ListUserHandler)
	api.GET("/user/:id", actionUser.ReadUserHandler)
	api.DELETE("/user/:id", actionUser.DeleteUserHandler)
	api.POST("/user", actionUser.CreateUserHandler)
	api.PUT("/user", actionUser.UpdateUserHandler)

	api.GET("/jenis_surat", actionJenisSurat.ListJenisSuratHandler)
	api.GET("/jenis_surat/:id", actionJenisSurat.ReadJenisSuratHandler)
	api.DELETE("/jenis_surat/:id", actionJenisSurat.DeleteJenisSuratHandler)
	api.POST("/jenis_surat", actionJenisSurat.CreateJenisSuratHandler)
	api.PUT("/jenis_surat", actionJenisSurat.UpdateUserHandler)

	e.Logger.Fatal(e.Start(":9000"))

	return e
}
