package controllers

import (
	"net/http"

	"github.com/DrZiMo/watch-tracker-api-golang/models"
	"github.com/gin-gonic/gin"
)

func GetAllWatches(c *gin.Context) {
	watches, err := models.GetWatches()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't fetch watches",
			"err":     err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      false,
		"message": "Fetched watches successfully",
		"watches": watches,
	})
}

func GetSingleWatch() {

}

func GetUserWatch() {

}

func CreateWatch(c *gin.Context) {
	var watch models.Watch
	err := c.ShouldBindJSON(&watch)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse requested data",
			"err":     err.Error(),
		})

		return
	}

	userId := c.GetInt64("userId")

	err = watch.Create(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't create watch",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":      true,
		"message": "Created watch successfully",
		"watch":   watch,
	})

}

func UpdateWatch() {

}

func DeleteWatch() {

}
