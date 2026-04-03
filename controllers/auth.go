package controllers

import (
	"net/http"

	"github.com/DrZiMo/watch-tracker-api-golang/models"
	"github.com/DrZiMo/watch-tracker-api-golang/utils"
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
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't parse requested data",
			"err":     err.Error(),
		})

		return
	}

	err = user.Login()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't login users",
			"err":     err.Error(),
		})

		return
	}

	token, err := utils.GenerateToke(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to generate token",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "User logged in successfully",
		"user":    user,
		"token":   token,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't parse requested data",
			"err":     err.Error(),
		})

		return
	}

	err = user.Create()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't create users",
			"err":     err.Error(),
		})

		return
	}

	token, err := utils.GenerateToke(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to generate token",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":      true,
		"message": "User created successfully",
		"user":    user,
		"token":   token,
	})
}
