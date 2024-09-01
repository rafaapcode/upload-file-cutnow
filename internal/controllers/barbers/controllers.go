package barbers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	aws_s3 "github.com/rafaapcode/upload-file-cutnow/pkg/aws"
	database_pkg "github.com/rafaapcode/upload-file-cutnow/pkg/mongo"
	controller_response "github.com/rafaapcode/upload-file-cutnow/types"
)

func BannerUpload(c echo.Context) error {
	var database database_pkg.Database

	id := c.FormValue("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, controller_response.Response{Status: false, Message: "Id é obrigatório !"})
	}

	database.HexId = id

	file, err := c.FormFile("file")
	if file.Size > int64(32897612) {
		return c.JSON(http.StatusNotAcceptable, controller_response.Response{Status: false, Message: "A imagem deve ter menos de 32MB"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado , tente novamente.", Error: err})
	}

	src, err := file.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao abrir o arquivo.", Error: err})
	}
	defer src.Close()

	client := database_pkg.Connect()
	database.Client = client
	defer database.Disconnect()

	filePath := fmt.Sprintf("barber/%s/banner-%s", id, file.Filename)
	aws_s3.UploadSingleFile("cutnow-images", filePath, src)

	_, err = database.UpdateBarberBanner(filePath)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao inserir o arquivo no Banco de dados.", Error: err})
	}

	return c.JSON(http.StatusCreated, controller_response.Response{Status: true, Message: "Banner uploaded with Successful !", Error: nil})
}

func FotoUpload(c echo.Context) error {
	var database database_pkg.Database
	id := c.FormValue("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, controller_response.Response{Status: false, Message: "Id é obrigatório !"})
	}
	database.HexId = id
	file, err := c.FormFile("file")
	if file.Size > int64(32897612) {
		return c.JSON(http.StatusNotAcceptable, controller_response.Response{Status: false, Message: "A imagem deve ter menos de 32MB"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado , tente novamente.", Error: err})
	}

	src, err := file.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao abrir o arquivo.", Error: err})
	}
	defer src.Close()

	client := database_pkg.Connect()
	database.Client = client
	defer database.Disconnect()

	filePath := fmt.Sprintf("barber/%s/foto-%s", id, file.Filename)
	aws_s3.UploadSingleFile("cutnow-images", filePath, src)

	_, err = database.UpdateBarberFoto(filePath)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao inserir o arquivo no Banco de dados.", Error: err})
	}

	return c.JSON(http.StatusCreated, controller_response.Response{Status: true, Message: "Foto uploaded with Successful !", Error: nil})
}

func PortfolioUpload(c echo.Context) error {
	var database database_pkg.Database
	id := c.FormValue("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, controller_response.Response{Status: false, Message: "Id é obrigatório !"})
	}
	database.HexId = id
	form, err := c.MultipartForm()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado , tente novamente.", Error: err})
	}

	if len(form.File) > 16 {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Você pode enviar no máximo 15 imagens", Error: err})
	}

	var filepaths []string

	client := database_pkg.Connect()
	database.Client = client
	defer database.Disconnect()

	for _, fileheaders := range form.File {
		file := fileheaders[0]

		if file.Size > int64(32897612) {
			return c.JSON(http.StatusNotAcceptable, controller_response.Response{Status: false, Message: "A imagem deve ter menos de 32MB"})
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao abrir o arquivo.", Error: err})
		}

		defer src.Close()
		filePath := fmt.Sprintf("barber/%s/potfolio-%s", id, file.Filename)
		filepaths = append(filepaths, filePath)
		go aws_s3.UploadMultipleFile("cutnow-images", filePath, src)
	}

	_, err = database.UpdateBarberPotfolio(filepaths)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao inserir o arquivo no Banco de dados.", Error: err})
	}

	return c.JSON(http.StatusCreated, controller_response.Response{Status: true, Message: "Portfolio photos uploaded with Successful !", Error: nil})
}
