package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
)

type (
	AdminController interface {
		RegisterAdmin()
		GetAllMandorForAdmin(ctx *gin.Context)
		GetDetailMandor(ctx *gin.Context)
		LoginAdmin(ctx *gin.Context)
		MeAdmin(ctx *gin.Context)
	}

	adminController struct {
		adminService services.AdminService
		jwtService services.JWTService
	}
)

func NewAdminController(as services.AdminService, jwtService services.JWTService) AdminController {
	return &adminController{
		adminService: as,
		jwtService: jwtService,
	}
}

func (ac *adminController) RegisterAdmin() {

}

func (ac *adminController) GetAllMandorForAdmin(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	adminID, err := ac.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	admin, err := ac.adminService.GetAdminByID(ctx.Request.Context(), adminID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
		return
	}

	if admin.Role != helpers.ADMIN {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", "Role Tidak Sesuai", utils.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
		return
	}

	mandors, err := ac.adminService.GetAllMandorForAdmin(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Mandor", mandors)
	ctx.JSON(http.StatusOK, res)
} 

func (ac *adminController) GetDetailMandor(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	adminID, err := ac.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	admin, err := ac.adminService.GetAdminByID(ctx.Request.Context(), adminID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
		return
	}

	if admin.Role != helpers.ADMIN {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", "Role Tidak Sesuai", utils.EmptyObj{})
		ctx.JSON(http.StatusForbidden, res)
		return
	}

	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	mandor, err := ac.adminService.GetDetailMandor(ctx.Request.Context(), uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Mandor", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Mandor", mandor)
	ctx.JSON(http.StatusOK, res)
}

func (ac *adminController) LoginAdmin(ctx *gin.Context) {
	var loginDTO dto.AdminLoginDTO
	if err := ctx.ShouldBind(&loginDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if check, err := ac.adminService.VerifyLogin(ctx.Request.Context(), loginDTO); !check {
		res := utils.BuildResponseFailed("Gagal Login", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	admin, err := ac.adminService.CheckAdminByEmail(ctx.Request.Context(), loginDTO.Email)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	if !admin.IsVerified {
		res := utils.BuildResponseFailed("Akun Belum Terverifikasi", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	if admin.Role != helpers.ADMIN {
		res := utils.BuildResponseFailed("Gagal Login", "Role Tidak Sesuai", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := ac.jwtService.GenerateToken(admin.ID, admin.Role)
	adminResponse := entities.Authorization {
		Token: token,
		Role: admin.Role,
	}

	res := utils.BuildResponseSuccess("Berhasil Login", adminResponse)
	ctx.JSON(http.StatusOK, res)
}

func (ac *adminController) MeAdmin(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	adminID, err := ac.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	admin, err := ac.adminService.GetAdminByID(ctx.Request.Context(), adminID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Admin", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Admin", admin)
	ctx.JSON(http.StatusOK, res)
}