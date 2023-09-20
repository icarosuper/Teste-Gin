package main

import (
	"Api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetRoutes(router)

	err := router.Run("localhost:3000")

	if err != nil {
		panic(err)
	}
}
