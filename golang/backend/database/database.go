package database

import (
	"fmt"
	"log"

	"github.com/mateus-schiavi/cardapio-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=capoeira dbname=cardapio_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = db.AutoMigrate(
		&models.Category{},
		&models.Dish{},
		&models.Rating{},
	)
	if err != nil {
		log.Fatal("Erro ao fazer AutoMigrate:", err)
	}

	DB = db
	fmt.Println("Banco de dados conectado e tabelas migradas com sucesso!")
}
