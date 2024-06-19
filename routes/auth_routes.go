package routes

import (
	"learn-gin-gorm/configs"
	"learn-gin-gorm/controllers"
	"learn-gin-gorm/middlewares"
	"learn-gin-gorm/repositories"
	"learn-gin-gorm/services"

	"github.com/gin-gonic/gin"
)

func AuthRouter() *gin.Engine {
	db := configs.InitDB()
	r := gin.Default()

	r.Use(middlewares.AuthMiddleware())

	authRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(authRepo)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(userService, authService)

	route := r.Group("/auth")
	{
		route.GET("/profile", authController.ProfileUser)
		route.POST("/register", authController.RegisterUser)
		route.POST("/login", authController.Login)
	}

	return r
}
