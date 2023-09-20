package routes

import (
	controllers "Api/controllers/user"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, db *sql.DB) {
	prefix := "/users"

	router.POST(prefix, func(c *gin.Context) {
		body := make([]byte, 0)

		_, err := c.Request.Body.Read(body)
		if err != nil {
			return
		}

		fmt.Println("username", string(body))

		err = controllers.CreateUser(string(body), db)

		if err != nil {

		}
	})
}
