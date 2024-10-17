package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	VideoID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DatabaseConfig struct {
	Host     string
	Database string
	User     string
	Pass     string
}

func loadConfig() *DatabaseConfig {
	// Load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database config
	db := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
	}
	return db
}

func GetDatabase() *gorm.DB {
	config := loadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Video{})

	return db
}
