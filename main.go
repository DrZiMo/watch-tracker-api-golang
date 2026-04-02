package main

import (
	"github.com/DrZiMo/watch-tracker-api-golang/db"
	"github.com/DrZiMo/watch-tracker-api-golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.AuthRoute(server)
	routes.WatchRoute(server)

	server.Run(":3000")
}
