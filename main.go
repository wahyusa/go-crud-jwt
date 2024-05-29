package main

import (
	"go-crud-jwt/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":" + os.Getenv("PORT"))
}
