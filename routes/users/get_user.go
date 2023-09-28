package routes

import (
	usercontroller "Api/controllers/user"
	"Api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type reqParams struct {
	id int `uri:"id" binding:"required,int"`
}

func GetUser(router *gin.Engine, db *gorm.DB) {
	router.GET("/users/:1", func(context *gin.Context) {
		var params reqParams

		if err := context.ShouldBindUri(&params); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err,
			})
			return
		}

		user, err := usercontroller.GetUser(params.id, db)

		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"body":    user,
		})
	})
}
