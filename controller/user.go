package controller

import (
	"net/http"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	MakeVerificationForgotPassword(ctx *gin.Context)
	ResendVerificationCode(ctx *gin.Context)
	ResendFailedLoginNotVerified(ctx *gin.Context)
	KodeOTPForgotPassword(ctx *gin.Context)
	SendForgotPassword(ctx *gin.Context)
	MeUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	jwtService  services.JWTService
	userService services.UserService
}

func NewUserController(us services.UserService, jwt services.JWTService) UserController {
	return &userController{
		jwtService:  jwt,
		userService: us,
	}
}

func (uc *userController) RegisterUser(ctx *gin.Context) {
	var user dto.UserCreateDTO
	if err := ctx.ShouldBind(&user); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if checkUser, _ := uc.userService.CheckUserEmail(ctx.Request.Context(), user.Email); checkUser {
		res := utils.BuildResponseFailed("Email Sudah Terdaftar", "failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if checkUser, _ := uc.userService.CheckUserUsername(ctx.Request.Context(), user.Username); checkUser {
		res := utils.BuildResponseFailed("Username Sudah Terdaftar", "failed", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.userService.RegisterUser(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) GetAllUser(ctx *gin.Context) {
	result, err := uc.userService.GetAllUser(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) MeUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	result, err := uc.userService.GetUserByID(ctx.Request.Context(), userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) ResendVerificationCode(ctx *gin.Context) {
	var resendVerificationCode dto.ResendVerificationCode
	if err := ctx.ShouldBind(&resendVerificationCode); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userVerification, err := uc.userService.ResendVerificationCode(ctx.Request.Context(), resendVerificationCode)
	if !userVerification {
		res := utils.BuildResponseFailed("Gagal Mengirim Ulang Kode Verifikasi", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return 
	}

	res := utils.BuildResponseSuccess("Berhasil Mengirim Ulang Kode Verifikasi", userVerification)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) LoginUser(ctx *gin.Context) {
	var userLoginDTO dto.UserLoginDTO
	var user entities.User
	if err := ctx.ShouldBind(&userLoginDTO); err != nil {
		response := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := uc.userService.VerifyLogin(ctx.Request.Context(), userLoginDTO)
	if !res {
		response := utils.BuildResponseFailed("Gagal Login", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if userLoginDTO.Email != "" {
		user, err = uc.userService.GetUserByEmail(ctx.Request.Context(), userLoginDTO.Email)
		if err != nil {
			response := utils.BuildResponseFailed("Gagal Login", err.Error(), utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else if userLoginDTO.Username != "" {
		user, err = uc.userService.GetUserByUsername(ctx.Request.Context(), userLoginDTO.Username)
		if err != nil {
			response := utils.BuildResponseFailed("Gagal Login", err.Error(), utils.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	if !user.IsVerified {
		response := utils.BuildResponseFailed("Gagal Login", "Email Belum Terverifikasi", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if user.Role != helpers.USER {
		response := utils.BuildResponseFailed("Gagal Login", "Role Tidak Sesuai", utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	token := uc.jwtService.GenerateToken(user.ID, user.Role)
	userResponse := entities.Authorization{
		Token: token,
		Role:  user.Role,
	}

	response := utils.BuildResponseSuccess("Berhasil Login", userResponse)
	ctx.JSON(http.StatusOK, response)
}

func (uc *userController) VerifyEmail(ctx *gin.Context) {
	var userVerificationDTO dto.UserVerificationDTO
	if err := ctx.ShouldBind(&userVerificationDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userVerification, err := uc.userService.VerifyEmail(ctx.Request.Context(), userVerificationDTO)
	if !userVerification {
		res := utils.BuildResponseFailed("Gagal Verifikasi Email", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Verifikasi Email", userVerification)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	var userDTO dto.UserUpdateDTO
	if err := ctx.ShouldBind(&userDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userDTO.ID = userID
	if err = uc.userService.UpdateUser(ctx.Request.Context(), userDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Update User", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Update User", userDTO)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := uc.jwtService.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err = uc.userService.DeleteUser(ctx.Request.Context(), userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus User", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menghapus User", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) ResendFailedLoginNotVerified(ctx *gin.Context) {
	var failedLoginDTO dto.FailedVerificationLoginDTO
	if err := ctx.ShouldBind(&failedLoginDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userVerification, err := uc.userService.ResendFailedLoginNotVerified(ctx.Request.Context(), failedLoginDTO.Email)
	if !userVerification {
		res := utils.BuildResponseFailed("Gagal Mengirim Ulang Kode Verifikasi", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return 
	}

	res := utils.BuildResponseSuccess("Berhasil Mengirim Ulang Kode Verifikasi", userVerification)
	ctx.JSON(http.StatusOK, res)
}

func(uc *userController) MakeVerificationForgotPassword(ctx *gin.Context) {
	var forgotPasswordReq dto.MakeVerificationForgotPasswordRequest
	if err := ctx.ShouldBind(&forgotPasswordReq); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	data, err := uc.userService.MakeVerificationForgotPassword(ctx.Request.Context(), forgotPasswordReq)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengirim Email", err.Error(), utils.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengirim Email", data)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) KodeOTPForgotPassword(ctx *gin.Context) {
	var req dto.KodeOTPForgotPasswordRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengirimkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := uc.userService.KodeOTPForgotPassword(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengirimkan Kode OTP", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengirimkan Kode OTP", data)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) SendForgotPassword(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengirimkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	data, err := uc.userService.SendForgotPassword(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengubah Password", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengganti Password", data)
	ctx.JSON(http.StatusOK, res)
}