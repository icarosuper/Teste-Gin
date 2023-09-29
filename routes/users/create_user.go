package routes

import (
	usercontroller "Api/controllers/user"
	"Api/dto"
	"Api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateUser(router *gin.Engine, db *gorm.DB) {
	router.POST("/users", func(context *gin.Context) {
		var user dto.UserDTO

		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err,
			})
			return
		}

		if err := usercontroller.CreateUser(user, db); err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Usuário criado com sucesso!",
		})
	})
}
