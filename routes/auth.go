package routes

import (
	"github.com/DrZiMo/watch-tracker-api-golang/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(server *gin.Engine) {
	server.GET("/auth", controllers.GetAllUsers)
	server.POST("/auth/login", controllers.LoginUser)
	server.POST("/auth/create", controllers.CreateUser)
}
