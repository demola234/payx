package routes

import (
	"go-service/payx/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
<<<<<<< HEAD
	incomingRoutes.GET("/account/:account_id", controllers.GetUserAccountDetails())
	incomingRoutes.GET("/card/:card_id", controllers.GetUserCardDetails())
	// incomingRoutes.GET("/users/:user_id", controllers.GetUser())
=======
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
>>>>>>> f475d850837453665936fdf77caeb7e2e97b1175
	// incomingRoutes.GET("/users/:user_id", controllers.EditUser())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser())
	incomingRoutes.POST("/users/login", controllers.Login())
	// incomingRoutes.POST("/users/profile", controllers.UploadProfileImage())
}
