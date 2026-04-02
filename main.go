package main

import (
	"log"

	"github.com/DrZiMo/watch-tracker-api-golang/db"
	"github.com/DrZiMo/watch-tracker-api-golang/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using defaults")
	}
}

func main() {
	db.InitDB()
	server := gin.Default()

	routes.AuthRoute(server)
	routes.WatchRoute(server)

	server.Run(":3000")
}
