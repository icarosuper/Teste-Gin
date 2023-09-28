package routes

import (
	"Api/routes/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	routes.UserRoutes(router, db)
}
