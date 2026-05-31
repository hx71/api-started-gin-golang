package handler

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/hx71/api-started-gin-golang/models"
)

func setupTestDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect test db: %v", err)
	}

	// migrate ulang supaya clean
	err = db.Migrator().DropTable(&models.Role{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Role{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
