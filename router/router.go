package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize routes with default configuration
	r := gin.Default()

	initializeRoutes(r)

	// Run ther server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
