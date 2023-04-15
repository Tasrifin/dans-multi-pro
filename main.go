package main

import (
	"dans-multi-pro/config"
	"dans-multi-pro/controllers"
	"dans-multi-pro/repositories"
	"dans-multi-pro/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	route := gin.Default()

	// user
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// route
	userRouter := route.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.Login)
	}

	route.Run(config.APP_PORT)
}
