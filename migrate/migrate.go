package main

import (
	"start-go/initializers"
	"start-go/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Shoes{})
}
