package routes

import "github.com/gin-gonic/gin"

// SetupRouter setup routes and route handlers
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/repos", getReposHandler)
	router.POST("/repos", updateReposHandler)
	return router
}
