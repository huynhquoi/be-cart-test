package models

import "gorm.io/gorm"

type Shoes struct {
	gorm.Model
	Image       string
	Name        string
	Description string
	Price       float64
	Color       string
}
