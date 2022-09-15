package routes

import (
	"go-service/payx/controllers"

	"github.com/gin-gonic/gin"
)

func WalletRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/account/:account_id", controllers.GetUserAccountDetailsByID())
	incomingRoutes.GET("/number/:account_number", controllers.GetUserAccountDetailsByNumber())
	incomingRoutes.GET("/card/:card_id", controllers.GetUserCardDetails())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.PATCH("/users/:user_id", controllers.UpdateUser())
	incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser())
}
