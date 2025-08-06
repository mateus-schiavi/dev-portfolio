package models

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	ImageURL    string
	CategoryID  uint
	Ratings     []Rating `gorm:"foreignKey:DishID"`
}

