package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	DishID uint
	Score int `gorm:"not null"`
	Comment string
}