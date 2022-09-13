package main

import (
	routers "go-service/payx/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	// router.Use(middleware.Authentication())
	routers.UserRoutes(router)
	// my part 
	// get all users
	// get one
	// deposit
	router.Run(":" + port)
}
