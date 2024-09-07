package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rafaapcode/upload-file-cutnow/internal/controllers/barbers"
	"github.com/rafaapcode/upload-file-cutnow/internal/controllers/barbershop"
)

func init() {
	os.Setenv("PORT", ":3002")
	os.Setenv("SECRET", "3d5af22c0142b5711b81dd51712d6454aa2e0870a5309128c0c77039b65fa94a")
	os.Setenv("MONGODB_URI", "mongodb+srv://rafaapcode:dXNcr6y312hhPo6V@cutnow.ald8nke.mongodb.net/cutnow?retryWrites=true&w=majority&appName=CutNow")
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// e.Use(middlewares.AuthMiddleware)

	e.GET("/",
		func(c echo.Context) error {
			return c.String(http.StatusOK, "Banner upload")
		})

	e.POST("/barbershop/logo", barbershop.LogoUpload)
	e.POST("/barbershop/banner", barbershop.BannerUpload)
	e.POST("/barbershop/structure", barbershop.StructureUpload)
	e.DELETE("/barbershop/structure/:index/:id", barbershop.DeleteStructImage)

	e.POST("/barber/foto", barbers.FotoUpload)
	e.POST("/barber/banner", barbers.BannerUpload)
	e.POST("/barber/portfolio", barbers.PortfolioUpload)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
