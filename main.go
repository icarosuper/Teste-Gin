package main

import (
	"Api/database"
	"Api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.GetDatabase()

	router := gin.Default()

	routes.SetRoutes(router, db)

	err = router.Run("localhost:3000")

	if err != nil {
		panic(err)
	}
}
