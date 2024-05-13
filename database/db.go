package database

import (
	"log"
	"os"

	"github.com/adrianramadhan/final-task-pbi-rakamin-fullstack-AdrianRamadhan/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Loading .env file")
	}

	dsn := os.Getenv("URL_DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Photo{}) 

	DB = db
}