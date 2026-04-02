package controllers

import (
	"net/http"

	"github.com/DrZiMo/watch-tracker-api-golang/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := models.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't fetch users",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Fetched users successfully",
		"users":   users,
	})
}

func LoginUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"users": "suli"})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"users": "suli"})
}
