package controllers

import (
	"net/http"
	"strconv"

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

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Fetched watches successfully",
		"watches": watches,
	})
}

func GetSingleWatch(c *gin.Context) {
	watchId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't get watch id",
			"err":     err.Error(),
		})

		return
	}

	watch, err := models.GetWatchByID(watchId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't fetch watch",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Fetched watch successfully",
		"watch":   watch,
	})

}

func GetUserWatch(c *gin.Context) {
	user_id := c.GetInt64("userId")
	watches, err := models.GetUserWatch(user_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't fetch user watch",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Fetched user watch successfully",
		"watches": watches,
	})
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

func UpdateWatch(c *gin.Context) {
	watchId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't get watch id",
			"err":     err.Error(),
		})

		return
	}

	var watch models.Watch
	err = c.ShouldBindJSON(&watch)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't parse requested data",
			"err":     err.Error(),
		})

		return
	}

	err = watch.Update(watchId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't update watch",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Watch updated successfully",
	})
}

func DeleteWatch(c *gin.Context) {
	watchId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Couldn't get watch id",
			"err":     err.Error(),
		})

		return
	}

	err = models.Delete(watchId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Couldn't delete watch",
			"err":     err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Watch deleted successfully",
	})
}
