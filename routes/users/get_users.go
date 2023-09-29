package routes

import (
	usercontroller "Api/controllers/user"
	"Api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUsers(router *gin.Engine, db *gorm.DB) {
	router.GET("/users", func(context *gin.Context) {
		users, err := usercontroller.GetUsers(db)

		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"body":    users,
		})
	})
}
