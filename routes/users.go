package routes

import (
	controllers "Api/controllers/user"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoutes(router *gin.Engine, db *sql.DB) {
	prefix := "/users"

	router.POST(prefix, func(c *gin.Context) {
		body := make([]byte, 0)

		_, err := c.Request.Body.Read(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		err = controllers.CreateUser(string(body), db)

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
