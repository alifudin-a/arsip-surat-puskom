package helper

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CustomReadableError(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s harus valid dan tidak boleh kosong",
					err.Field())
			case "numeric":
				report.Message = fmt.Sprintf("%s harus berupa angka",
					err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s tidak valid",
					err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}

			break
		}
	}
	c.JSON(report.Code, report)
}
