package main

import (
	routers "go-service/payx/server/routes"
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

	router.Run(":" + port)
}
