package routes

import (
	AuthRoutes "Api/routes/auth"
	UserRoutes "Api/routes/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	AuthRoutes.AuthRoutes(router, db)
	UserRoutes.UserRoutes(router, db)
}
