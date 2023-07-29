package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/gin-gonic/gin"
)

func Image(route *gin.Engine, ImageController controller.ImageController) {
	routes := route.Group("/api/image")
	{
		routes.POST("", ImageController.UploadImage)
		routes.GET("/get/storage/:path/:dirname/:filename", ImageController.GetImage)
	}
}