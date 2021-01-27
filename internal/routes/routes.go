package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/repos", getReposHandler)
	router.POST("/repos", updateReposHandler)
	return router
}
