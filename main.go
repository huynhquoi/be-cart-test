package main

import (
	"start-go/controllers"
	"start-go/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{"http://localhost:3000/"}

	corsConfig.AllowCredentials = true

	corsConfig.AddAllowMethods("OPTIONS")

	router.Use(cors.New(corsConfig))

	router.GET("/shoes", controllers.ShoesIndex)
	router.GET("/shoes/:id", controllers.ShoesById)
	router.POST("/shoes", controllers.CreateShoes)

	router.Run()
}
