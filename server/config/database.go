package config

import (
	"fmt"
	"log"

	"github.com/sagakyuuu/mahasiswa/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=ssint password=123123 dbname=mahasiswa_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database Connected Failed")

	}

	if err := db.AutoMigrate(&models.User{}, &models.Profile{}); err != nil {
		log.Fatal("Error:", err)
	}
	DB = db

	fmt.Println("Database Connected Successfuly")
}
