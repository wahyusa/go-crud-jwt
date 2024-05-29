package routes

import (
	"go-crud-jwt/controllers"
	"go-crud-jwt/models"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	models.ConnectDatabase()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(controllers.AuthMiddleware())
	{
		authorized.GET("/user", controllers.GetUser)
	}
}
