package routes

import (
	controllers "Api/controllers/user"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(router *gin.Engine, db *sql.DB) {
	router.GET("/users", func(context *gin.Context) {
		controllers.GetUsers(db)

		context.JSON(http.StatusOK, gin.H{
			"message": "blz",
		})
	})
}
