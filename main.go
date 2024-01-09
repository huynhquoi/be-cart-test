package main

import (
	"start-go/controllers"
	"start-go/initializers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	router.GET("/shoes", controllers.ShoesIndex)
	router.GET("/shoes/:id", controllers.ShoesById)
	router.POST("/shoes", controllers.CreateShoes)

	router.Run()
}
