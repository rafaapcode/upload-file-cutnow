package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/upload-file-cutnow/pkg/jwt"
)

type Response struct {
	Message string `json:"message"`
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, Response{Message: "Authorization header is required"})
		}

		_, err := jwt.ValidatingToken(authHeader)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Response{Message: "Token invalid"})
		}

		return next(c)
	}
}
