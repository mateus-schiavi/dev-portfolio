package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateus-schiavi/cardapio-backend/database"
	"github.com/mateus-schiavi/cardapio-backend/models"
)

func GetDishes(c *gin.Context) {
	var dishes []models.Dish
	database.DB.Preload("Ratings").Find(&dishes)
	c.JSON(http.StatusOK, dishes)
}

func CreateDish(c *gin.Context) {
	var dish models.Dish
	if err := c.ShouldBindJSON(&dish); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&dish)
	c.JSON(http.StatusCreated, dish)
}

func GetDishByID(c *gin.Context) {
	id := c.Param("id")
	var dish models.Dish
	if err := database.DB.Preload("Ratings").First(&dish, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prato não encontrado"})
		return
	}
	c.JSON(http.StatusOK, dish)
}
