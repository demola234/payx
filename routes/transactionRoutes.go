package routes

import (
	controller "go-service/payx/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/transaction/deposit", controller.Deposit())
	incomingRoutes.GET("/transaction/verify/:ref", controller.Verify())
	incomingRoutes.GET("/transaction/user", controller.GetUserTransaction())
}
