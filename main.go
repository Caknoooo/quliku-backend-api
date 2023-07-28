package main

import (
	"log"
	"os"

	"github.com/Caknoooo/golang-clean_template/config"
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/controller/seederController"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/migrations"
	"github.com/Caknoooo/golang-clean_template/repository"
	"github.com/Caknoooo/golang-clean_template/repository/seederRepository"
	"github.com/Caknoooo/golang-clean_template/routes"
	"github.com/Caknoooo/golang-clean_template/routes/seeder"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/Caknoooo/golang-clean_template/services/seederServices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db                         *gorm.DB                              = config.SetUpDatabaseConnection()
		jwtService                 services.JWTService                   = services.NewJWTService()
		userVerificationRepository repository.UserVerificationRepository = repository.NewUserVerificationRepository(db)
		userRepository             repository.UserRepository             = repository.NewUserRepository(db)
		userService                services.UserService                  = services.NewUserService(userRepository, userVerificationRepository)
		userController             controller.UserController             = controller.NewUserController(userService, jwtService)
		imageRepository            repository.ImageRepository            = repository.NewImageRepository(db)
		imageService               services.ImageService                 = services.NewImageService(imageRepository)
		imageController						controller.ImageController            = controller.NewImageController(imageService, jwtService)
		mandorRepository 				 repository.MandorRepository           = repository.NewMandorRepository(db)
		mandorVerificationRepository = repository.NewMandorVerificationRepository(db)
		mandorService						 services.MandorService                = services.NewMandorService(mandorRepository, mandorVerificationRepository)
		mandorController				 controller.MandorController           = controller.NewMandorController(mandorService, jwtService)
		adminRepository 					repository.AdminRepository            = repository.NewAdminRepository(db)
		adminService 						 services.AdminService                 = services.NewAdminService(adminRepository)
		adminController 				 controller.AdminController            = controller.NewAdminController(adminService, jwtService)
	)

	// Seeder
	var (
		listBankRepository seederRepository.ListBankRepository = seederRepository.NewListBankRepository(db)
		listBankService    seederServices.ListBankService      = seederServices.NewListBankService(listBankRepository)
		listBankController seederController.ListBankController = seederController.NewListBankController(listBankService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	routes.User(server, userController, jwtService)
	routes.Image(server, imageController)
	routes.Mandor(server, mandorController, jwtService)
	routes.Admin(server, adminController, jwtService)

	// Seeder Routes
	seeder.ListBank(server, listBankController)

	if err := migrations.Seeder(db); err != nil {
		log.Fatalf("error migration seeder: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run(":" + port)
}
