package controller

import (
	"net/http"
	"strconv"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/utils"
	"github.com/gin-gonic/gin"
)

type (
	ProjectUserController interface {
		CreateProjectUser(ctx *gin.Context)
		GetAllProjectUser(ctx *gin.Context)
		GetDetailProjectUser(ctx *gin.Context)
		ChangeStatusProjectUser(ctx *gin.Context)
		GetAllProjectUserByUserId(ctx *gin.Context)
		GetDetailProjectUserById(ctx *gin.Context)
	}

	projectUserController struct {
		s   services.ProjectUserService
		jwt services.JWTService
	}
)

func NewProjectUserController(s services.ProjectUserService, jwt services.JWTService) *projectUserController {
	return &projectUserController{
		s:   s,
		jwt: jwt,
	}
}

func (c *projectUserController) CreateProjectUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := c.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	var req dto.CreateProjectRequest

	ctx.PostFormArray("detail_category")
	ctx.PostFormArray("proof_of_damage")
	ctx.PostFormArray("type_of_craftsman")

	req.Category = ctx.Request.PostForm.Get("category")
	req.DetailDescription = ctx.Request.PostForm.Get("detail_description")
	req.Alamat = ctx.Request.PostForm.Get("alamat")
	req.DateStart = ctx.Request.PostForm.Get("date_start")
	req.DateEnd = ctx.Request.PostForm.Get("date_end")
	req.Estimation = ctx.Request.PostForm.Get("estimation")
	req.TotalPrice, _ = strconv.Atoi(ctx.Request.PostForm.Get("total_price"))

	for i := 0; ; i++ {
		name := ctx.Request.PostForm.Get("detail_category[" + strconv.Itoa(i) + "][name]")
		if name == "" {
			break
		}

		var detailCategory dto.DetailCategoryRequest
		detailCategory.Name = name

		isActive, _ := strconv.ParseBool(ctx.Request.PostForm.Get("detail_category[" + strconv.Itoa(i) + "][is_active]"))

		detailCategory.IsActive = isActive
		req.DetailCategoryRequest = append(req.DetailCategoryRequest, detailCategory)
	}

	for i := 0; ; i++ {
		photo, err := ctx.FormFile("proof_of_damage[" + strconv.Itoa(i) + "][photo]")
		if err != nil {
			break
		}

		if photo == nil {
			break
		}

		var proofOfDamage dto.ProofOfDamageRequest
		proofOfDamage.Photo = photo
		req.ProofOfDamageRequest = append(req.ProofOfDamageRequest, proofOfDamage)
	}

	for i := 0; ; i++ {
		var typeOfCraftsman dto.TypeOfCraftsmanRequest

		duration := ctx.Request.PostForm.Get("type_of_craftsman[" + strconv.Itoa(i) + "][duration]")
		if duration == "" {
			break
		}
		isHalfDay, _ := strconv.ParseBool(ctx.Request.PostForm.Get("type_of_craftsman[" + strconv.Itoa(i) + "][is_half_day]"))
		price, _ := strconv.Atoi(ctx.Request.PostForm.Get("type_of_craftsman[" + strconv.Itoa(i) + "][price]"))
		craftType := ctx.Request.PostForm.Get("type_of_craftsman[" + strconv.Itoa(i) + "][type]")
		if isHalfDay {
			typeOfCraftsman.IsHalfDay = true
			typeOfCraftsman.IsFullDay = false
		} else {
			typeOfCraftsman.IsHalfDay = false
			typeOfCraftsman.IsFullDay = true
		}

		typeOfCraftsman.Duration = duration
		typeOfCraftsman.Price = price
		typeOfCraftsman.Type = craftType

		req.TypeOfCraftsmanRequest = append(req.TypeOfCraftsmanRequest, typeOfCraftsman)
	}

	projectUser, err := c.s.CreateProjectUser(ctx, req, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Membuat Project", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Membuat Project", projectUser)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectUserController) GetAllProjectUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := c.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	projectUser, err := c.s.GetAllProjectUserByAdmin(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Project", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Project", projectUser)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectUserController) GetDetailProjectUser(ctx *gin.Context) {
	id := ctx.Param("id")

	token := ctx.MustGet("token").(string)
	userID, err := c.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	projectUser, err := c.s.GetProjectUserByIdByAdmin(ctx, userID, id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Project", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Project", projectUser)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectUserController) ChangeStatusProjectUser(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	adminId, err := c.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	var req dto.ChangeStatusProjectUserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	projectUser, err := c.s.ChangeStatusProjectUser(ctx, adminId, req)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengubah Status Project", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengubah Status Project", projectUser)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectUserController) GetAllProjectUserByUserId(ctx *gin.Context) {
	token := ctx.MustGet("token").(string)
	userID, err := c.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	projectUser, err := c.s.GetAllProjectUserByUserId(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Project", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Project", projectUser)
	ctx.JSON(http.StatusOK, res)
}

func (c *projectUserController) GetDetailProjectUserById(ctx *gin.Context) {
	projectId := ctx.Param("project_id")

	token := ctx.MustGet("token").(string)
	userID, err := c.jwt.GetIDByToken(token)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	projectUser, err := c.s.GetDetailProjectUserById(ctx, projectId, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Project", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Project", projectUser)
	ctx.JSON(http.StatusOK, res)
}