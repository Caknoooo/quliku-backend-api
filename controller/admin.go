package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/helpers"
)

type (
	AdminController interface {
		RegisterAdmin()
		LoginAdmin(ctx *gin.Context)
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