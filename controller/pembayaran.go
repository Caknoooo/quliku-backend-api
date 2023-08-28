package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type (
	PembayaranController interface {
		Create(ctx *gin.Context)
		GetPembayaranById(ctx *gin.Context)
		GetAllPembayaranByUserId(ctx *gin.Context)
		GetAllPembayaran(ctx *gin.Context)
	}

	pembayaranController struct {
		ps  services.PembayaranService
		jwt services.JWTService
	}
)

func NewPembayaranController(ps services.PembayaranService, jwt services.JWTService) PembayaranController {
	return &pembayaranController{
		ps:  ps,
		jwt: jwt,
	}
}

func (pc *pembayaranController) Create(ctx *gin.Context) {
	projectId := ctx.Param("project_id")

	token := ctx.MustGet("token").(string)
	userId, err := pc.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	var req dto.PembayaranRequest
	req.Name = ctx.PostForm("name")
	req.AccountNumber = ctx.Request.PostForm.Get("account_number")
	req.BankName = ctx.Request.PostForm.Get("bank_name")
	req.TotalPrice, _ = strconv.Atoi(ctx.Request.PostForm.Get("total_price"))
	req.PaymentPhoto, _ = ctx.FormFile("payment_photo")

	fmt.Println("Name:", req.Name)
	fmt.Println("AccountNumber:", req.AccountNumber	)
	fmt.Println("BankName:", req.BankName)
	fmt.Println("TotalPrice:", req.TotalPrice)
	fmt.Println("PaymentPhoto:", req.PaymentPhoto.Filename)

	pembayaran, err := pc.ps.Create(ctx, req, projectId, userId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Membuat Pembayaran", pembayaran)
	ctx.JSON(http.StatusCreated, res)
}

func (pc *pembayaranController) GetPembayaranById(ctx *gin.Context) {
	id := ctx.Param("id")

	token := ctx.MustGet("token").(string)
	adminId, err := pc.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	pembayaran, err := pc.ps.GetPembayaranById(ctx, adminId, id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Detail Pembayaran", pembayaran)
	ctx.JSON(http.StatusOK, res)
}

func (pc *pembayaranController) GetAllPembayaranByUserId(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userId, err := pc.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	pembayaran, err := pc.ps.GetAllPembayaranByUserId(ctx, userId.String())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Detail Pembayaran", pembayaran)
	ctx.JSON(http.StatusOK, res)
}

func (pc *pembayaranController) GetAllPembayaran(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	adminId, err := pc.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	pembayaran, err := pc.ps.GetAllPembayaran(ctx, adminId)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Detail Pembayaran", pembayaran)
	ctx.JSON(http.StatusOK, res)
}