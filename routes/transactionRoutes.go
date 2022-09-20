package routes

import (
	"go-service/payx/controllers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/transaction/deposit", controllers.Deposit())
}

// func WalletRoutes(incomingRoutes *gin.Engine){
// 	incomingRoutes.Get("/transaction/user", controllers.GetUsersTransactions())
	
// }