package routes

import (
	usercontroller "Api/controllers/user"
	"Api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type GetUserParams struct {
	id int `uri:"id" binding:"required, int"`
}

func GetUser(router *gin.Engine, db *gorm.DB) {
	router.GET("/users/:1", func(context *gin.Context) {
		var params GetUserParams

		if err := context.ShouldBindUri(&params); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err,
			})
			return
		}

		user, err := usercontroller.GetUserById(params.id, db)

		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"user":    user,
		})
	})
}
