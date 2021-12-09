package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConfig() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Can not retreive environment variables")
	}
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Can not connect to your database")
	} else {
		fmt.Println("Successfully connected to the database.")
	}

	return db
}
