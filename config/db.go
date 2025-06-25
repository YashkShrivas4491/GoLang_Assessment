package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"makerble-crudapi/models"
)

var DB *gorm.DB
var Err error

func InitDB() {
	// Load .env file
	 godotenv.Load()
	

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	DB, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{}, &models.Patient{})
}
