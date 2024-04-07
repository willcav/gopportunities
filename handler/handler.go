package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}

func ListOpeningsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
