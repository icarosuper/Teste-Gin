package routes

import (
	usercontroller "Api/controllers/user"
	"Api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUsers(router *gin.Engine, db *gorm.DB) {
	router.GET("/users", func(context *gin.Context) {
		users, err := usercontroller.GetUsers(db)

		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		//fmt.Println(users[0].ID)
		//fmt.Println(users[0].Name)
		//fmt.Println(users[0].Email)
		//fmt.Println(users[0].Birthdate)
		//fmt.Println(users[0].Role)
		//fmt.Println(users[0].CreatedAt)
		//fmt.Println(users[0].UpdatedAt)
		//fmt.Println(users[0].ActivatedAt)

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"body":    users,
		})
	})
}
