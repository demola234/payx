package main

import (
	"go-service/payx/middleware"
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
	router.Use(middleware.Authentication())
	routers.WalletRoutes(router)
<<<<<<< HEAD
=======
	routers.TransactionRoutes(router)
>>>>>>> 03c6fec9e38190dca9703b7139b9428066fddd72
	router.Run(":" + port)
}
