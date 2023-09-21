package routes

import (
	"Api/routes/users"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"nig":     "lesgoo",
		})
	})

	routes.UserRoutes(router, db)
}
