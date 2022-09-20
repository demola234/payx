package main

import (
	// "go-service/payx/middleware"
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
	routers.UserRoutes(router)
	// router.Use(middleware.Authentication())
	routers.WalletRoutes(router)
	routers.TransactionRoutes(router)
	router.Run(":" + port)
}
