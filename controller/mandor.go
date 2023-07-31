package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type MandorController interface {
	RegisterMandorStart(ctx *gin.Context)
	RegisterMandorEnd(ctx *gin.Context)
	LoginMandor(ctx *gin.Context)
	MeMandor(ctx *gin.Context)
	UpdateMandor(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	ResendVerificationCode(ctx *gin.Context)
	ResendFailedLoginNotVerified(ctx *gin.Context)
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

	if !mandorx.IsVerified {
		res := utils.BuildResponseFailed("Gagal Login", "Akun Belum Terverifikasi", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	if mandorx.Role != helpers.MANDOR {
		res := utils.BuildResponseFailed("Gagal Login", "Role Tidak Sesuai", utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
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

func (mc *mandorController) MeMandor(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	ID, err := mc.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	mandor, err := mc.mandorService.GetMandorByMandorID(ctx.Request.Context(), ID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Mandor", mandor)
	ctx.JSON(http.StatusOK, res)
}

func (mc *mandorController) VerifyEmail(ctx *gin.Context) {
	var mandorVerificationDTO dto.MandorVerificationDTO
	if err := ctx.ShouldBind(&mandorVerificationDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	mandorVerification, err := mc.mandorService.VerifyEmail(ctx.Request.Context(), mandorVerificationDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Verifikasi Email", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Verifikasi Email", mandorVerification)
	ctx.JSON(http.StatusOK, res)
}

func (mc *mandorController) ResendVerificationCode(ctx *gin.Context) {
	var mandorVerificationDTO dto.ResendMandorVerificationCodeDTO
	if err := ctx.ShouldBind(&mandorVerificationDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	mandorVerification, err := mc.mandorService.ResendVerificationCode(ctx.Request.Context(), mandorVerificationDTO)
	if !mandorVerification {
		res := utils.BuildResponseFailed("Gagal Mengirim Ulang Kode Verifikasi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengirim Ulang Kode Verifikasi", mandorVerification)
	ctx.JSON(http.StatusOK, res)
}

func (mc *mandorController) ResendFailedLoginNotVerified(ctx *gin.Context) {
	var failedLoginDTO dto.FailedMandorVerificationLoginDTO
	if err := ctx.ShouldBind(&failedLoginDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	failedLogin, err := mc.mandorService.ResendFailedLoginNotVerified(ctx.Request.Context(), failedLoginDTO.Email)
	if !failedLogin {
		res := utils.BuildResponseFailed("Gagal Mengirim Ulang Kode Verifikasi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengirim Ulang Kode Verifikasi", failedLogin)
	ctx.JSON(http.StatusOK, res)
}

func (mc *mandorController) UpdateMandor(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	ID, err := mc.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var mandorUpdateDTO dto.MandorUpdateDTO
	if err := ctx.ShouldBind(&mandorUpdateDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	mandorUpdateDTO.ID = ID
	data, err := mc.mandorService.UpdateMandor(ctx.Request.Context(), mandorUpdateDTO)
	if err != nil{
		res := utils.BuildResponseFailed("Gagal Update Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Update Mandor", data)
	ctx.JSON(http.StatusOK, res)
}