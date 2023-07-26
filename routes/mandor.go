package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/gin-gonic/gin"
)

func Mandor(route *gin.Engine, MandorController controller.MandorController) {
	routes := route.Group("/api/mandor")
	{
		routes.POST("", MandorController.RegisterMandorStart)
		routes.POST("/next", MandorController.RegisterMandorEnd)
		routes.POST("/login", MandorController.LoginMandor)
	}
}