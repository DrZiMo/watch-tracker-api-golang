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

func CreateWatch() {

}

func UpdateWatch() {

}

func DeleteWatch() {

}
