package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	SignUp(router, db)
	SignIn(router, db)
}
