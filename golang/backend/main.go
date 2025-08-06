package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mateus-schiavi/cardapio-backend/controllers"
	"github.com/mateus-schiavi/cardapio-backend/database"
)

func main() {
	// Conecta ao banco e migra as tabelas
	database.Connect()

	// Cria o servidor Gin
	r := gin.Default()

	// Rotas para categorias
	r.GET("/categories", controllers.GetCategories)
	r.POST("/categories", controllers.CreateCategory)

	// Rotas para pratos
	r.GET("/dishes", controllers.GetDishes)
	r.GET("/dishes/:id", controllers.GetDishByID)
	r.POST("/dishes", controllers.CreateDish)

	// Rota para avaliação
	r.POST("/ratings", controllers.CreateRating)

	// Roda o servidor na porta 8080
	r.Run(":8080")
}
