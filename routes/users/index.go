package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB) {
	GetUsers(router, db)
	GetUser(router, db)
	CreateUser(router, db)
}
