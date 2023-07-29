package controller

import (
	// "fmt"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	// "github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
)

type ImageController interface {
	UploadImage(ctx *gin.Context)
	GetImage(ctx *gin.Context)
}

type imageController struct {
	imageService services.ImageService
	jwtService   services.JWTService
}

func NewImageController(is services.ImageService, js services.JWTService) ImageController {
	return &imageController{
		imageService: is,
		jwtService:   js,
	}
}

func (ic *imageController) UploadImage(ctx *gin.Context) {
	// var ImageForm dto.ImageUploadDTO
	// var ImageCreate dto.ImageCreateDTO
	file, err := ctx.FormFile("image_form")
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	bytes, err := utils.IsBase64(*file) 
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	// fmt.Println(file)

	err = utils.SaveImage(bytes, "images", "UMUM", file.Filename)
	if err != nil {
		response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	fmt.Print("Sukses")
}

func (ic *imageController) GetImage(ctx *gin.Context) {
	// var image string = ""
	path := ctx.Param("path")
	dirName := ctx.Param("dirname")
	file := ctx.Param("filename")
	imagePath := "storage" + "/" + path + "/" + dirName + "/" + file

	_, err := os.Stat(imagePath)
	if err != nil {
		if os.IsNotExist(err) {
			response := utils.BuildResponseFailed("Failed to process request", err.Error(), utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	ctx.File(imagePath)
}