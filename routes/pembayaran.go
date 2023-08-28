package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Pembayaran(route *gin.Engine, PembayaranController controller.PembayaranController, JWTService services.JWTService) {
	routes := route.Group("/api/pembayaran") 
	{
		routes.POST("/:project_id", middleware.Authenticate(JWTService), PembayaranController.Create)
		routes.GET("/:id", middleware.Authenticate(JWTService), PembayaranController.GetPembayaranById)
		routes.GET("", middleware.Authenticate(JWTService), PembayaranController.GetAllPembayaran)
		routes.GET("/user", middleware.Authenticate(JWTService), PembayaranController.GetAllPembayaranByUserId)
	}
}