package database

import (
	"final_project_golang/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	if isDebugMode() {
		runMigrations()
	}
}

func GetDB() *gorm.DB {

	if isDebugMode() {
		return db.Debug()
	}

	return db
}

func isDebugMode() bool {
	debugModeStr := os.Getenv("DEBUG_MODE")
	debugMode, err := strconv.ParseBool(debugModeStr)
	if err != nil {
		return false
	}
	return debugMode
}

func runMigrations() {
	err := db.AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}
}
