package seeder

import (
	"github.com/Caknoooo/golang-clean_template/controller/seederController"
	"github.com/gin-gonic/gin"
)

func ListBank(route *gin.Engine, ListbankController seederController.ListBankController) {
	routes := route.Group("/api/seeder") 
	{
		routes.GET("/list_bank", ListbankController.GetAllListBank)
		routes.GET("/list_bank/:id", ListbankController.GetBankByID)
	}
}