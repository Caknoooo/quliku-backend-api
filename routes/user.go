package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, UserController controller.UserController, jwtService services.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", UserController.RegisterUser)
		routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		routes.POST("/login", UserController.LoginUser)
		routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
		routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)

		// Verifikasi akun
		routes.POST("/send_verification_forgot_password", UserController.MakeVerificationForgotPassword)
		routes.POST("/send_verification", UserController.ResendVerificationCode)
		routes.POST("/verify", UserController.VerifyEmail)
		routes.POST("/failed_login/verify", UserController.ResendFailedLoginNotVerified)
	}
}
