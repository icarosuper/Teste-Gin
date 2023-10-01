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

func SignUp(router *gin.Engine, db *gorm.DB) {
	router.POST("/signup", func(context *gin.Context) {
		var userInfo dto.SignUpDto

		if err := context.ShouldBindJSON(&userInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err,
			})
			return
		}

		if password, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), 10); err != nil {
			utils.SendUnknownError(err, context)
		} else {
			userInfo.Password = string(password)
		}

		createdUser, err := usercontroller.CreateUser(userInfo, db)
		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		token, err := authcontroller.CreateJwt(createdUser.Id, time.Hour*24)
		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success":   true,
			"message":   "Usuário criado com sucesso!",
			"authToken": token,
		})
	})
}
