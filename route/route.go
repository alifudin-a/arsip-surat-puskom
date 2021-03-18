package route

import (
	"net/http"

	action "github.com/alifudin-a/arsip-surat-puskom/action/jabatan-struktural"
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
	api.GET("/jabatan_struktural", action.ListJabatanStrukturalHandler)
	api.GET("/jabatan_struktural/:id", action.ReadJabatanStrukturalHandler)
	api.DELETE("/jabatan_struktural/:id", action.DeleteJabatanStrukturalHandler)

	e.Logger.Fatal(e.Start(":9000"))

	return e
}
