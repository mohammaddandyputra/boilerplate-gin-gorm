package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// err = database.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }

	DB = database

	return DB
}
