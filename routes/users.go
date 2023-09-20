package routes

import (
	controllers "Api/controllers/user"
	"Api/entities"
	"Api/utils"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoutes(router *gin.Engine, db *sql.DB) {
	prefix := "/users"

	router.POST(prefix, func(c *gin.Context) {
		user, err := utils.ParseRequestBody[entities.User](c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		err = controllers.CreateUser(user, db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"success": true,
			"message": "Usuário criado com sucesso!",
		})
	})
}
