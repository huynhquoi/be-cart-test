package main

import (
	"start-go/controllers"
	"start-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/shoes", controllers.ShoesIndex)
	router.GET("/shoes/:id", controllers.ShoesById)
	router.POST("/shoes", controllers.CreateShoes)

	router.Run()
}
