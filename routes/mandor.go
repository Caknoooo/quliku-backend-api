package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Mandor(route *gin.Engine, MandorController controller.MandorController, jwtService services.JWTService) {
	routes := route.Group("/api/mandor")
	{
		routes.POST("", MandorController.RegisterMandorStart)
		routes.POST("/next", MandorController.RegisterMandorEnd)
		routes.POST("/login", MandorController.LoginMandor)
		routes.GET("/me", middleware.Authenticate(jwtService), MandorController.MeMandor)
	}
}