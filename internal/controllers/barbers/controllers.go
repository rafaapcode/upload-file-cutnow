package barbers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BannerUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Banner upload")
}

func FotoUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Foto upload")
}

func PortfolioUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Portfolio upload")
}
