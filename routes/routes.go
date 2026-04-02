package routes

import "github.com/gin-gonic/gin"

func RegisterServer(server *gin.Engine) {
	// Auth
	server.GET("/auth")
	server.POST("/auth/login")
	server.POST("/auth/create")

	// Watch
	server.GET("/watches")
	server.GET("/watch/:id")
	server.GET("/watch/user/:id")
	server.POST("/watch/")
	server.PUT("/watch/:id")
	server.DELETE("/watch/:id")
}
