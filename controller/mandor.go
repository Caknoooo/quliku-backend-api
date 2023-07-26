package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type MandorController interface {
	RegisterMandorStart(ctx *gin.Context)
	RegisterMandorEnd(ctx *gin.Context)
	LoginMandor(ctx *gin.Context)
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

func (mc *mandorController) LoginMandor(ctx *gin.Context) {
	var loginDTO dto.MandorLoginDTO
	var mandorx entities.Mandor
	if err := ctx.ShouldBind(&loginDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	mandor, err := mc.mandorService.VerifyLogin(ctx.Request.Context(), loginDTO)
	if !mandor {
		res := utils.BuildResponseFailed("Gagal Login Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	mandorx, err = mc.mandorService.CheckMandorByEmail(ctx.Request.Context(), loginDTO.Email)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Login Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := mc.jwtService.GenerateToken(mandorx.ID, mandorx.Role)
	mandorResponse := entities.Authorization{
		Token: token,
		Role: mandorx.Role,
	}

	res := utils.BuildResponseSuccess("Berhasil Login Mandor", mandorResponse)
	ctx.JSON(http.StatusOK, res)
}