package seederController

import (
	"net/http"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/services/seederServices"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type (
	ListBankController interface {
		GetAllListBank(ctx *gin.Context)
		GetBankByID(ctx *gin.Context)
	}

	listBankController struct {
		listBankService seederServices.ListBankService
	}
)

func NewListBankController(listBankService seederServices.ListBankService) ListBankController {
	return &listBankController{
		listBankService: listBankService,
	}
}

func (l *listBankController) GetAllListBank(ctx *gin.Context) {
	listBank, err := l.listBankService.GetAllListBank(ctx.Request.Context())
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Bank", listBank)
	ctx.JSON(http.StatusOK, res)
}

func (l *listBankController) GetBankByID(ctx *gin.Context) {
	id := ctx.Param("id")
	bankID, err := strconv.Atoi(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse ID", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := l.listBankService.GetBankByID(ctx.Request.Context(), bankID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Bank", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Bank", result)
	ctx.JSON(http.StatusOK, res)
}