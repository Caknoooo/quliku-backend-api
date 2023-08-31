package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func ProjectUser(route *gin.Engine, projectUserController controller.ProjectUserController, jwtService services.JWTService) {
	routes := route.Group("/api/project_user")
	{
		routes.POST("", middleware.Authenticate(jwtService), projectUserController.CreateProjectUser)
		routes.GET("", middleware.Authenticate(jwtService), projectUserController.GetAllProjectUserByUserId)
		routes.GET("/:project_id", middleware.Authenticate(jwtService), projectUserController.GetDetailProjectUserById)
	}
}