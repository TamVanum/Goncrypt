package config

import (
	"log"

	"github.com/tamVanum/goncrypt/internal/user/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	log.Println("✅ Connected to database")

	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed migrations:", err)
	}
}
