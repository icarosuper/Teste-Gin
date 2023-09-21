package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, db *sql.DB) {
	GetUsers(router, db)
	CreateUser(router, db)
}
