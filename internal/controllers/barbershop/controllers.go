package barbershop

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	aws_s3 "github.com/rafaapcode/upload-file-cutnow/pkg/aws"
	controller_response "github.com/rafaapcode/upload-file-cutnow/types"
)

func BannerUpload(c echo.Context) error {
	id := c.FormValue("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, controller_response.Response{Status: false, Message: "Id é obrigatório !"})
	}

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
	filePath := fmt.Sprintf("barbershop/%s/banner-%s", id, file.Filename)
	aws_s3.UploadSingleFile("cutnow-images", filePath, src)

	return c.JSON(http.StatusCreated, controller_response.Response{Status: true, Message: "Banner uploaded with Successful !", Error: nil})
}

func LogoUpload(c echo.Context) error {
	id := c.FormValue("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, controller_response.Response{Status: false, Message: "Id é obrigatório !"})
	}

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
	filePath := fmt.Sprintf("barbershop/%s/logo-%s", id, file.Filename)
	aws_s3.UploadSingleFile("cutnow-images", filePath, src)
	return c.JSON(http.StatusCreated, controller_response.Response{Status: true, Message: "Logo uploaded with Successful !", Error: nil})
}

func StructureUpload(c echo.Context) error {
	id := c.FormValue("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, controller_response.Response{Status: false, Message: "Id é obrigatório !"})
	}

	form, err := c.MultipartForm()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado , tente novamente.", Error: err})
	}

	for _, fileheaders := range form.File {

		file := fileheaders[0]

		if len(fileheaders) > 6 {
			return c.JSON(http.StatusNotAcceptable, controller_response.Response{Status: false, Message: "Você pode enviar no máximo 6 imagens.", Error: err})
		}
		if file.Size > int64(32897612) {
			return c.JSON(http.StatusNotAcceptable, controller_response.Response{Status: false, Message: "A imagem deve ter menos de 32MB"})
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, controller_response.Response{Status: false, Message: "Algo deu errado ao abrir o arquivo.", Error: err})
		}

		defer src.Close()
		filePath := fmt.Sprintf("barbershop/%s/structure-%s", id, file.Filename)
		go aws_s3.UploadMultipleFile("cutnow-images", filePath, src)
	}
	return c.JSON(http.StatusCreated, controller_response.Response{Status: true, Message: "Structure photos uploaded with Successful!", Error: nil})
}
