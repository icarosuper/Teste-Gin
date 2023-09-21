package routes

import (
	controllers "Api/controllers/user"
	"Api/dto"
	"Api/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(router *gin.Engine, db *sql.DB) {
	router.POST("/users", func(context *gin.Context) {
		reqUser, err := utils.ParseRequestBody[dto.GetUserFromReqDTO](context)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		err = controllers.CreateUser(dto.ParseReqUser(reqUser), db)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Usuário criado com sucesso!",
		})
	})
}
