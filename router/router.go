package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize routes with default configuration
	r := gin.Default()

	// Set a GET route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the api
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
