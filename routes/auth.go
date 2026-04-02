package routes

import (
	"github.com/DrZiMo/watch-tracker-api-golang/controllers"
	"github.com/DrZiMo/watch-tracker-api-golang/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoute(server *gin.Engine) {
	server.GET("/auth", middleware.Authenticate, controllers.GetAllUsers)
	server.POST("/auth/login", controllers.LoginUser)
	server.POST("/auth/create", controllers.CreateUser)
}
