package routes

import (
	"go-service/payx/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/account/:account_id", controllers.GetUserAccountDetailsByID())
	incomingRoutes.GET("/account/:account_number", controllers.GetUserAccountDetailsByNumber())
	incomingRoutes.GET("/card/:card_id", controllers.GetUserCardDetails())
	// incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	// incomingRoutes.GET("/users/:user_id", controllers.EditUser())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser())
	incomingRoutes.POST("/users/login", controllers.Login())
	// incomingRoutes.POST("/users/profile", controllers.UploadProfileImage())
}
