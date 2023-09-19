package configs

import (
	"fmt"
	"gin-todo/pkg/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable that holds the database connection
func DB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	Database_str := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(Database_str), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Todo{})

	return db
}
