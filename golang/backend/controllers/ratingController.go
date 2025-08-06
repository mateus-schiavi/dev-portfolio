package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateus-schiavi/cardapio-backend/database"
	"github.com/mateus-schiavi/cardapio-backend/models"
)

func CreateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&rating)
	c.JSON(http.StatusCreated, rating)
}
