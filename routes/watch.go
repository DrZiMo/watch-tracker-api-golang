package routes

import (
	"github.com/DrZiMo/watch-tracker-api-golang/controllers"
	"github.com/DrZiMo/watch-tracker-api-golang/middleware"
	"github.com/gin-gonic/gin"
)

func WatchRoute(server *gin.Engine) {
	server.GET("/watch", middleware.Authenticate, controllers.GetAllWatches)
	server.GET("/watch/:id")
	server.GET("/watch/user/:id")
	server.POST("/watch/")
	server.PUT("/watch/:id")
	server.DELETE("/watch/:id")
}
