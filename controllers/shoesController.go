package controllers

import (
	"start-go/initializers"
	"start-go/models"

	"github.com/gin-gonic/gin"
)

func CreateShoes(c *gin.Context) {

	var body struct {
		Image       string
		Name        string
		Description string
		Price       float64
		Color       string
	}

	c.Bind(&body)

	shoes := models.Shoes{
		Image:       body.Image,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Color:       body.Color}
	result := initializers.DB.Create(&shoes)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"shoes": shoes,
	})
}

func ShoesIndex(c *gin.Context) {
	var shoes []models.Shoes
	initializers.DB.Find(&shoes)

	c.JSON(200, gin.H{
		"shoes": shoes,
	})
}

func ShoesById(c *gin.Context) {
	id := c.Param("id")

	var shoes models.Shoes
	initializers.DB.First(&shoes, id)

	c.JSON(200, gin.H{
		"shoes": shoes,
	})
}
