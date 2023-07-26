package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type MandorController interface {
	RegisterMandorStart(ctx *gin.Context)
	RegisterMandorEnd(ctx *gin.Context)
}

type mandorController struct {
	mandorService services.MandorService
	jwtService    services.JWTService
}

func NewMandorController(ms services.MandorService, jwt services.JWTService) MandorController {
	return &mandorController{
		mandorService: ms,
		jwtService:    jwt,
	}
}

func (mc *mandorController) RegisterMandorStart(ctx *gin.Context) {
	var mandorCreateDTO dto.MandorCreateDTO
	if err := ctx.ShouldBind(&mandorCreateDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, err := mc.mandorService.RegisterMandorStart(ctx.Request.Context(), mandorCreateDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Mandor", data)
	ctx.JSON(http.StatusOK, res)
}

func(mc *mandorController) RegisterMandorEnd(ctx *gin.Context) {
	var mandorCreateDTO dto.MandorNextDTO

	if err := ctx.ShouldBind(&mandorCreateDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, err := mc.mandorService.RegisterMandorEnd(ctx.Request.Context(), mandorCreateDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Mandor", data)
	ctx.JSON(http.StatusOK, res)
}