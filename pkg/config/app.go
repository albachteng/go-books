package config

import (
	"log"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	DB_URI := os.Getenv("POSTGRES_URL")
	// NTS replace second param with sql connection
	d, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
