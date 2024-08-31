package barbershop

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func BannerUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Banner upload")
}

func LogoUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Logo upload")
}

func StructureUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Structure upload")
}
