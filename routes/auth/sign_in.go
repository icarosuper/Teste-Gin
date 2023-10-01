package routes

import (
	authcontroller "Api/controllers/auth"
	usercontroller "Api/controllers/user"
	"Api/dto"
	"Api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func SignIn(router *gin.Engine, db *gorm.DB) {
	router.POST("/signin", func(context *gin.Context) {
		var userInfo dto.SignInDto

		if err := context.ShouldBindJSON(&userInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err,
			})
			return
		}

		user, err := usercontroller.GetUserByEmail(userInfo.Email, db)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "email não encontrado",
			})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "senha incorreta",
			})
			return
		}

		token, err := authcontroller.CreateJwt(user.Id, time.Hour*24)
		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success":   true,
			"message":   "logado com sucesso",
			"authToken": token,
		})
	})
}
