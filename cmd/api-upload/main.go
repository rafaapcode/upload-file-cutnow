package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rafaapcode/upload-file-cutnow/internal/controllers/barbers"
	"github.com/rafaapcode/upload-file-cutnow/internal/controllers/barbershop"
)

func init() {
	os.Setenv("PORT", ":3001")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/barbershop/logo", barbershop.LogoUpload)
	e.POST("/barbershop/banner", barbershop.BannerUpload)
	e.POST("/barbershop/structure", barbershop.StructureUpload)

	e.POST("/barber/foto", barbers.FotoUpload)
	e.POST("/barber/banner", barbers.BannerUpload)
	e.POST("/barber/portfolio", barbers.PortfolioUpload)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
