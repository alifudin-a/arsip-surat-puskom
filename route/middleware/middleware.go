package middleware

import (
	"fmt"
	"os"

	codes "github.com/alifudin-a/arsip-surat-puskom/domain/helper"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

func Validation() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = ValidationJWT(c)
			if err != nil {
				return err
			}

			err = ValidationKey(c)
			if err != nil {
				return err
			}

			return h(c)
		}
	}
}

func ValidationKey(c echo.Context) (err error) {
	key := c.Request().Header.Get("api-key")
	if key != os.Getenv("ASU-API-KEY") {
		return status.Errorf(
			codes.InvalidAPIKey,
			fmt.Sprintf("%v", "API Key tidak cocok!"),
		)
	}
	return
}

func ValidationJWT(c echo.Context) (err error) {
	tokenString := c.Request().Header.Get("user-token")

	// jika token kosong
	if tokenString == "" {
		return status.Error(codes.InvalidToken, "Token Wajib Diisi!")
	}

	token, err := jwt.Parse(tokenString, jwtToken)

	if token.Valid {
		return
	} else if validationError, ok := err.(*jwt.ValidationError); ok {
		if validationError.Errors&jwt.ValidationErrorMalformed != 0 {
			// jika bukan token
			return status.Error(codes.InvalidToken, "Token Wajib Diisi!")
		} else if validationError.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// token sudah expired atau belum aktif
			return status.Error(codes.InvalidToken, "Token Sudah Kadaluarsa!")
		} else {
			// token tidak bisa diproses
			return status.Errorf(codes.InvalidToken, fmt.Sprintf("%v", "Token tidak bisa diverifikasi!"))
		}
	} else {
		// token tidak bisa diproses
		return status.Errorf(codes.InvalidToken, fmt.Sprintf("%v", "Token tidak bisa diverifikasi!"))
	}
}

func jwtToken(token *jwt.Token) (interface{}, error) {
	return []byte(fmt.Sprintf(os.Getenv("JWT_SECRET"))), nil
}
