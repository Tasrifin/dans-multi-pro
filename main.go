package main

import (
	"dans-multi-pro/config"
	"dans-multi-pro/controllers"
	"dans-multi-pro/middlewares"
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

	// job
	jobRepo := repositories.NewJobRepo(db)
	jobService := services.NewJobService(jobRepo)
	jobController := controllers.NewJobController(jobService)

	// route
	route.POST("/register", userController.UserRegister)
	route.POST("/login", userController.Login)

	jobRouter := route.Group("/job")
	{
		jobRouter.Use(middlewares.Auth())
		jobRouter.GET("/list", jobController.GetJobList)
		jobRouter.GET("/detail/:detailID", jobController.GetJobDetail)
	}

	route.Run(config.APP_PORT)
}
