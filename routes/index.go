package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"nig":     "lesgoo",
		})
	})
}
