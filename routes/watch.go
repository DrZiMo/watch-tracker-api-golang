package routes

import (
	"github.com/DrZiMo/watch-tracker-api-golang/controllers"
	"github.com/DrZiMo/watch-tracker-api-golang/middleware"
	"github.com/gin-gonic/gin"
)

func WatchRoute(server *gin.Engine) {
	server.GET("/watch", middleware.Authenticate, controllers.GetAllWatches)
	server.GET("/watch/:id", middleware.Authenticate, controllers.GetSingleWatch)
	server.GET("/watch/user", middleware.Authenticate, controllers.GetUserWatch)
	server.POST("/watch", middleware.Authenticate, controllers.CreateWatch)
	server.PUT("/watch/:id", middleware.Authenticate, controllers.UpdateWatch)
	server.DELETE("/watch/:id", middleware.Authenticate, controllers.DeleteWatch)
}
