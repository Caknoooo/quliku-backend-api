package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, AdminController controller.AdminController, projectUser controller.ProjectUserController, jwtService services.JWTService) {
	routes := route.Group("/api/admin")
	{
		routes.POST("/login", AdminController.LoginAdmin)
		routes.GET("/me", middleware.Authenticate(jwtService), AdminController.MeAdmin)

		routes.GET("/get/mandor", middleware.Authenticate(jwtService), AdminController.GetAllMandorForAdmin)
		routes.GET("/get/mandor/:id", middleware.Authenticate(jwtService), AdminController.GetDetailMandor)
		routes.POST("/update/mandor", middleware.Authenticate(jwtService), AdminController.ChangeStatusMandor)

		routes.GET("/get/project_user", middleware.Authenticate(jwtService), projectUser.GetAllProjectUser)
		routes.GET("/get/project_user/:id", middleware.Authenticate(jwtService), projectUser.GetDetailProjectUser)
		routes.PATCH("/update/project_user", middleware.Authenticate(jwtService), projectUser.ChangeStatusProjectUser)
	}
}